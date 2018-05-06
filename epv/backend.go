                                         
                                                
  
                                                                                  
                                                                              
                                                                    
                                      
  
                                                                             
                                                                 
                                                               
                                                      
  
                                                                           
                                                                                  

                                                
package epv

import (
	"errors"
	"fmt"
	"math/big"
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/epvchain/go-epvchain/act"
	"github.com/epvchain/go-epvchain/public"
	"github.com/epvchain/go-epvchain/public/hexutil"
	"github.com/epvchain/go-epvchain/agreement"
	"github.com/epvchain/go-epvchain/agreement/epvdpos"
	"github.com/epvchain/go-epvchain/agreement/epvhash"
	"github.com/epvchain/go-epvchain/kernel"
	"github.com/epvchain/go-epvchain/kernel/bloombits"
	"github.com/epvchain/go-epvchain/kernel/types"
	"github.com/epvchain/go-epvchain/kernel/vm"
	"github.com/epvchain/go-epvchain/epv/downloader"
	"github.com/epvchain/go-epvchain/epv/filters"
	"github.com/epvchain/go-epvchain/epv/gasprice"
	"github.com/epvchain/go-epvchain/data"
	"github.com/epvchain/go-epvchain/notice"
	"github.com/epvchain/go-epvchain/local/epvapi"
	"github.com/epvchain/go-epvchain/book"
	"github.com/epvchain/go-epvchain/work"
	"github.com/epvchain/go-epvchain/point"
	"github.com/epvchain/go-epvchain/peer"
	"github.com/epvchain/go-epvchain/content"
	"github.com/epvchain/go-epvchain/process"
	"github.com/epvchain/go-epvchain/remote"
)

type LesServer interface {
	Start(srvr *p2p.Server)
	Stop()
	Protocols() []p2p.Protocol
	SetBloomBitsIndexer(bbIndexer *core.ChainIndexer)
}

                                                      
type EPVchain struct {
	config      *Config
	chainConfig *params.ChainConfig

	                                        
	shutdownChan  chan bool                                             
	stopDbUpgrade func() error                                        

	           
	txPool          *core.TxPool
	blockchain      *core.BlockChain
	protocolManager *ProtocolManager
	lesServer       LesServer

	                
	chainDb epvdb.Database                        

	eventMux       *event.TypeMux
	engine         consensus.Engine
	accountManager *accounts.Manager

	bloomRequests chan chan *bloombits.Retrieval                                                   
	bloomIndexer  *core.ChainIndexer                                                            

	ApiBackend *EPVApiBackend

	miner     *miner.Miner
	gasPrice  *big.Int
	epvcbase common.Address

	networkId     uint64
	netRPCService *epvapi.PublicNetAPI

	lock sync.RWMutex                                                              
}

func (s *EPVchain) AddLesServer(ls LesServer) {
	s.lesServer = ls
	ls.SetBloomBitsIndexer(s.bloomIndexer)
}

                                                   
                                                
func New(ctx *node.ServiceContext, config *Config) (*EPVchain, error) {
	if config.SyncMode == downloader.LightSync {
		return nil, errors.New("can't run epv.EPVchain in light sync mode, use les.LightEPVchain")
	}
	if !config.SyncMode.IsValid() {
		return nil, fmt.Errorf("invalid sync mode %d", config.SyncMode)
	}
	chainDb, err := CreateDB(ctx, config, "chaindata")
	if err != nil {
		return nil, err
	}
	stopDbUpgrade := upgradeDeduplicateData(chainDb)
	chainConfig, genesisHash, genesisErr := core.SetupGenesisBlock(chainDb, config.Genesis)
	if _, ok := genesisErr.(*params.ConfigCompatError); genesisErr != nil && !ok {
		return nil, genesisErr
	}
	log.Info("Initialised chain configuration", "config", chainConfig)

	epv := &EPVchain{
		config:         config,
		chainDb:        chainDb,
		chainConfig:    chainConfig,
		eventMux:       ctx.EventMux,
		accountManager: ctx.AccountManager,
		engine:         CreateConsensusEngine(ctx, &config.EPVhash, chainConfig, chainDb),
		shutdownChan:   make(chan bool),
		stopDbUpgrade:  stopDbUpgrade,
		networkId:      config.NetworkId,
		gasPrice:       config.GasPrice,
		epvcbase:      config.EPVCbase,
		bloomRequests:  make(chan chan *bloombits.Retrieval),
		bloomIndexer:   NewBloomIndexer(chainDb, params.BloomBitsBlocks),
	}

	log.Info("Initialising EPVchain protocol", "versions", ProtocolVersions, "network", config.NetworkId)

	if !config.SkipBcVersionCheck {
		bcVersion := core.GetBlockChainVersion(chainDb)
		if bcVersion != core.BlockChainVersion && bcVersion != 0 {
			return nil, fmt.Errorf("Blockchain DB version mismatch (%d / %d). Run gepv upgradedb.\n", bcVersion, core.BlockChainVersion)
		}
		core.WriteBlockChainVersion(chainDb, core.BlockChainVersion)
	}
	var (
		vmConfig    = vm.Config{EnablePreimageRecording: config.EnablePreimageRecording}
		cacheConfig = &core.CacheConfig{Disabled: config.NoPruning, TrieNodeLimit: config.TrieCache, TrieTimeLimit: config.TrieTimeout}
	)
	epv.blockchain, err = core.NewBlockChain(chainDb, cacheConfig, epv.chainConfig, epv.engine, vmConfig)
	if err != nil {
		return nil, err
	}
	                                                              
	if compat, ok := genesisErr.(*params.ConfigCompatError); ok {
		log.Warn("Rewinding chain to upgrade configuration", "err", compat)
		epv.blockchain.SetHead(compat.RewindTo)
		core.WriteChainConfig(chainDb, genesisHash, chainConfig)
	}
	epv.bloomIndexer.Start(epv.blockchain)

	if config.TxPool.Journal != "" {
		config.TxPool.Journal = ctx.ResolvePath(config.TxPool.Journal)
	}
	epv.txPool = core.NewTxPool(config.TxPool, epv.chainConfig, epv.blockchain)

	if epv.protocolManager, err = NewProtocolManager(epv.chainConfig, config.SyncMode, config.NetworkId, epv.eventMux, epv.txPool, epv.engine, epv.blockchain, chainDb); err != nil {
		return nil, err
	}
	epv.miner = miner.New(epv, epv.chainConfig, epv.EventMux(), epv.engine)
	epv.miner.SetExtra(makeExtraData(config.ExtraData))

	epv.ApiBackend = &EPVApiBackend{epv, nil}
	gpoParams := config.GPO
	if gpoParams.Default == nil {
		gpoParams.Default = config.GasPrice
	}
	epv.ApiBackend.gpo = gasprice.NewOracle(epv.ApiBackend, gpoParams)

	return epv, nil
}

func makeExtraData(extra []byte) []byte {
	if len(extra) == 0 {
		                           
		extra, _ = rlp.EncodeToBytes([]interface{}{
			uint(params.VersionMajor<<16 | params.VersionMinor<<8 | params.VersionPatch),
			"gepv",
			runtime.Version(),
			runtime.GOOS,
		})
	}
	if uint64(len(extra)) > params.MaximumExtraDataSize {
		log.Warn("Miner extra data exceed limit", "extra", hexutil.Bytes(extra), "limit", params.MaximumExtraDataSize)
		extra = nil
	}
	return extra
}

                                       
func CreateDB(ctx *node.ServiceContext, config *Config, name string) (epvdb.Database, error) {
	db, err := ctx.OpenDatabase(name, config.DatabaseCache, config.DatabaseHandles)
	if err != nil {
		return nil, err
	}
	if db, ok := db.(*epvdb.LDBDatabase); ok {
		db.Meter("epv/db/chaindata/")
	}
	return db, nil
}

                                                                                                       
func CreateConsensusEngine(ctx *node.ServiceContext, config *epvhash.Config, chainConfig *params.ChainConfig, db epvdb.Database) consensus.Engine {
	if chainConfig.DPos != nil {
		return epvdpos.New(chainConfig.DPos, db)
	}
	                                 
	switch {
	case config.PowMode == epvhash.ModeFake:
		log.Warn("EPVhash used in fake mode")
		return epvhash.NewFaker()
	case config.PowMode == epvhash.ModeTest:
		log.Warn("EPVhash used in test mode")
		return epvhash.NewTester()
	case config.PowMode == epvhash.ModeShared:
		log.Warn("EPVhash used in shared mode")
		return epvhash.NewShared()
	default:
		engine := epvhash.New(epvhash.Config{
			CacheDir:       ctx.ResolvePath(config.CacheDir),
			CachesInMem:    config.CachesInMem,
			CachesOnDisk:   config.CachesOnDisk,
			DatasetDir:     config.DatasetDir,
			DatasetsInMem:  config.DatasetsInMem,
			DatasetsOnDisk: config.DatasetsOnDisk,
		})
		engine.SetThreads(-1)                      
		return engine
	}
}

                                                                           
                                                                            
func (s *EPVchain) APIs() []rpc.API {
	apis := epvapi.GetAPIs(s.ApiBackend)

	                                                             
	apis = append(apis, s.engine.APIs(s.BlockChain())...)

	                                       
	return append(apis, []rpc.API{
		{
			Namespace: "epv",
			Version:   "1.0",
			Service:   NewPublicEPVchainAPI(s),
			Public:    true,
		}, {
			Namespace: "epv",
			Version:   "1.0",
			Service:   NewPublicMinerAPI(s),
			Public:    true,
		}, {
			Namespace: "epv",
			Version:   "1.0",
			Service:   downloader.NewPublicDownloaderAPI(s.protocolManager.downloader, s.eventMux),
			Public:    true,
		}, {
			Namespace: "miner",
			Version:   "1.0",
			Service:   NewPrivateMinerAPI(s),
			Public:    false,
		}, {
			Namespace: "epv",
			Version:   "1.0",
			Service:   filters.NewPublicFilterAPI(s.ApiBackend, false),
			Public:    true,
		}, {
			Namespace: "admin",
			Version:   "1.0",
			Service:   NewPrivateAdminAPI(s),
		}, {
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewPublicDebugAPI(s),
			Public:    true,
		}, {
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewPrivateDebugAPI(s.chainConfig, s),
		}, {
			Namespace: "net",
			Version:   "1.0",
			Service:   s.netRPCService,
			Public:    true,
		},
	}...)
}

func (s *EPVchain) ResetWithGenesisBlock(gb *types.Block) {
	s.blockchain.ResetWithGenesisBlock(gb)
}

func (s *EPVchain) EPVCbase() (eb common.Address, err error) {
	s.lock.RLock()
	epvcbase := s.epvcbase
	s.lock.RUnlock()

	if epvcbase != (common.Address{}) {
		return epvcbase, nil
	}
	if wallets := s.AccountManager().Wallets(); len(wallets) > 0 {
		if accounts := wallets[0].Accounts(); len(accounts) > 0 {
			epvcbase := accounts[0].Address

			s.lock.Lock()
			s.epvcbase = epvcbase
			s.lock.Unlock()

			log.Info("EPVCbase automatically configured", "address", epvcbase)
			return epvcbase, nil
		}
	}
	return common.Address{}, fmt.Errorf("epvcbase must be explicitly specified")
}

                                                                  
func (self *EPVchain) SetEPVCbase(epvcbase common.Address) {
	self.lock.Lock()
	self.epvcbase = epvcbase
	self.lock.Unlock()

	self.miner.SetEPVCbase(epvcbase)
}

func (s *EPVchain) StartMining(local bool) error {
	eb, err := s.EPVCbase()
	if err != nil {
		log.Error("Cannot start mining without epvcbase", "err", err)
		return fmt.Errorf("epvcbase missing: %v", err)
	}
	if dpos, ok := s.engine.(*epvdpos.DPos); ok {
		wallet, err := s.accountManager.Find(accounts.Account{Address: eb})
		if wallet == nil || err != nil {
			log.Error("EPVCbase account unavailable locally", "err", err)
			return fmt.Errorf("signer missing: %v", err)
		}
		dpos.Authorize(eb, wallet.SignHash)
	}
	if local {
		                                                                             
		                                                                               
		                                                                            
		                                                                   
		atomic.StoreUint32(&s.protocolManager.acceptTxs, 1)
	}
	go s.miner.Start(eb)
	return nil
}

func (s *EPVchain) StopMining()         { s.miner.Stop() }
func (s *EPVchain) IsMining() bool      { return s.miner.Mining() }
func (s *EPVchain) Miner() *miner.Miner { return s.miner }

func (s *EPVchain) AccountManager() *accounts.Manager  { return s.accountManager }
func (s *EPVchain) BlockChain() *core.BlockChain       { return s.blockchain }
func (s *EPVchain) TxPool() *core.TxPool               { return s.txPool }
func (s *EPVchain) EventMux() *event.TypeMux           { return s.eventMux }
func (s *EPVchain) Engine() consensus.Engine           { return s.engine }
func (s *EPVchain) ChainDb() epvdb.Database            { return s.chainDb }
func (s *EPVchain) IsListening() bool                  { return true }                    
func (s *EPVchain) EPVVersion() int                    { return int(s.protocolManager.SubProtocols[0].Version) }
func (s *EPVchain) NetVersion() uint64                 { return s.networkId }
func (s *EPVchain) Downloader() *downloader.Downloader { return s.protocolManager.downloader }

                                                                            
                              
func (s *EPVchain) Protocols() []p2p.Protocol {
	if s.lesServer == nil {
		return s.protocolManager.SubProtocols
	}
	return append(s.protocolManager.SubProtocols, s.lesServer.Protocols()...)
}

                                                                                
                                    
func (s *EPVchain) Start(srvr *p2p.Server) error {
	                                            
	s.startBloomHandlers()

	                        
	s.netRPCService = epvapi.NewPublicNetAPI(srvr, s.NetVersion())

	                                                          
	maxPeers := srvr.MaxPeers
	if s.config.LightServ > 0 {
		if s.config.LightPeers >= srvr.MaxPeers {
			return fmt.Errorf("invalid peer config: light peer count (%d) >= total peer count (%d)", s.config.LightPeers, srvr.MaxPeers)
		}
		maxPeers -= s.config.LightPeers
	}
	                                                               
	s.protocolManager.Start(maxPeers)
	if s.lesServer != nil {
		s.lesServer.Start(srvr)
	}
	return nil
}

                                                                                
                     
func (s *EPVchain) Stop() error {
	if s.stopDbUpgrade != nil {
		s.stopDbUpgrade()
	}
	s.bloomIndexer.Close()
	s.blockchain.Stop()
	s.protocolManager.Stop()
	if s.lesServer != nil {
		s.lesServer.Stop()
	}
	s.txPool.Stop()
	s.miner.Stop()
	s.eventMux.Stop()

	s.chainDb.Close()
	close(s.shutdownChan)

	return nil
}
