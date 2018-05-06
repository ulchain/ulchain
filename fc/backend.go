                                         
                                                
  
                                                                                  
                                                                              
                                                                    
                                      
  
                                                                             
                                                                 
                                                               
                                                      
  
                                                                           
                                                                                  

                                                         
package les

import (
	"fmt"
	"sync"
	"time"

	"github.com/epvchain/go-epvchain/act"
	"github.com/epvchain/go-epvchain/public"
	"github.com/epvchain/go-epvchain/public/hexutil"
	"github.com/epvchain/go-epvchain/agreement"
	"github.com/epvchain/go-epvchain/kernel"
	"github.com/epvchain/go-epvchain/kernel/bloombits"
	"github.com/epvchain/go-epvchain/kernel/types"
	"github.com/epvchain/go-epvchain/epv"
	"github.com/epvchain/go-epvchain/epv/downloader"
	"github.com/epvchain/go-epvchain/epv/filters"
	"github.com/epvchain/go-epvchain/epv/gasprice"
	"github.com/epvchain/go-epvchain/data"
	"github.com/epvchain/go-epvchain/notice"
	"github.com/epvchain/go-epvchain/local/epvapi"
	"github.com/epvchain/go-epvchain/simple"
	"github.com/epvchain/go-epvchain/book"
	"github.com/epvchain/go-epvchain/point"
	"github.com/epvchain/go-epvchain/peer"
	"github.com/epvchain/go-epvchain/peer/discv5"
	"github.com/epvchain/go-epvchain/content"
	rpc "github.com/epvchain/go-epvchain/remote"
)

type LightEPVchain struct {
	config *epv.Config

	odr         *LesOdr
	relay       *LesTxRelay
	chainConfig *params.ChainConfig
	                                        
	shutdownChan chan bool
	           
	peers           *peerSet
	txPool          *light.TxPool
	blockchain      *light.LightChain
	protocolManager *ProtocolManager
	serverPool      *serverPool
	reqDist         *requestDistributor
	retriever       *retrieveManager
	                
	chainDb epvdb.Database                        

	bloomRequests                              chan chan *bloombits.Retrieval                                                   
	bloomIndexer, chtIndexer, bloomTrieIndexer *core.ChainIndexer

	ApiBackend *LesApiBackend

	eventMux       *event.TypeMux
	engine         consensus.Engine
	accountManager *accounts.Manager

	networkId     uint64
	netRPCService *epvapi.PublicNetAPI

	wg sync.WaitGroup
}

func New(ctx *node.ServiceContext, config *epv.Config) (*LightEPVchain, error) {
	chainDb, err := epv.CreateDB(ctx, config, "lightchaindata")
	if err != nil {
		return nil, err
	}
	chainConfig, genesisHash, genesisErr := core.SetupGenesisBlock(chainDb, config.Genesis)
	if _, isCompat := genesisErr.(*params.ConfigCompatError); genesisErr != nil && !isCompat {
		return nil, genesisErr
	}
	log.Info("Initialised chain configuration", "config", chainConfig)

	peers := newPeerSet()
	quitSync := make(chan struct{})

	lepv := &LightEPVchain{
		config:           config,
		chainConfig:      chainConfig,
		chainDb:          chainDb,
		eventMux:         ctx.EventMux,
		peers:            peers,
		reqDist:          newRequestDistributor(peers, quitSync),
		accountManager:   ctx.AccountManager,
		engine:           epv.CreateConsensusEngine(ctx, &config.EPVhash, chainConfig, chainDb),
		shutdownChan:     make(chan bool),
		networkId:        config.NetworkId,
		bloomRequests:    make(chan chan *bloombits.Retrieval),
		bloomIndexer:     epv.NewBloomIndexer(chainDb, light.BloomTrieFrequency),
		chtIndexer:       light.NewChtIndexer(chainDb, true),
		bloomTrieIndexer: light.NewBloomTrieIndexer(chainDb, true),
	}

	lepv.relay = NewLesTxRelay(peers, lepv.reqDist)
	lepv.serverPool = newServerPool(chainDb, quitSync, &lepv.wg)
	lepv.retriever = newRetrieveManager(peers, lepv.reqDist, lepv.serverPool)
	lepv.odr = NewLesOdr(chainDb, lepv.chtIndexer, lepv.bloomTrieIndexer, lepv.bloomIndexer, lepv.retriever)
	if lepv.blockchain, err = light.NewLightChain(lepv.odr, lepv.chainConfig, lepv.engine); err != nil {
		return nil, err
	}
	lepv.bloomIndexer.Start(lepv.blockchain)
	                                                              
	if compat, ok := genesisErr.(*params.ConfigCompatError); ok {
		log.Warn("Rewinding chain to upgrade configuration", "err", compat)
		lepv.blockchain.SetHead(compat.RewindTo)
		core.WriteChainConfig(chainDb, genesisHash, chainConfig)
	}

	lepv.txPool = light.NewTxPool(lepv.chainConfig, lepv.blockchain, lepv.relay)
	if lepv.protocolManager, err = NewProtocolManager(lepv.chainConfig, true, ClientProtocolVersions, config.NetworkId, lepv.eventMux, lepv.engine, lepv.peers, lepv.blockchain, nil, chainDb, lepv.odr, lepv.relay, quitSync, &lepv.wg); err != nil {
		return nil, err
	}
	lepv.ApiBackend = &LesApiBackend{lepv, nil}
	gpoParams := config.GPO
	if gpoParams.Default == nil {
		gpoParams.Default = config.GasPrice
	}
	lepv.ApiBackend.gpo = gasprice.NewOracle(lepv.ApiBackend, gpoParams)
	return lepv, nil
}

func lesTopic(genesisHash common.Hash, protocolVersion uint) discv5.Topic {
	var name string
	switch protocolVersion {
	case lpv1:
		name = "LES"
	case lpv2:
		name = "LES2"
	default:
		panic(nil)
	}
	return discv5.Topic(name + "@" + common.Bytes2Hex(genesisHash.Bytes()[0:8]))
}

type LightDummyAPI struct{}

                                                              
func (s *LightDummyAPI) EPVCbase() (common.Address, error) {
	return common.Address{}, fmt.Errorf("not supported")
}

                                                                                   
func (s *LightDummyAPI) Coinbase() (common.Address, error) {
	return common.Address{}, fmt.Errorf("not supported")
}

                                    
func (s *LightDummyAPI) Hashrate() hexutil.Uint {
	return 0
}

                                                                 
func (s *LightDummyAPI) Mining() bool {
	return false
}

                                                                           
                                                                            
func (s *LightEPVchain) APIs() []rpc.API {
	return append(epvapi.GetAPIs(s.ApiBackend), []rpc.API{
		{
			Namespace: "epv",
			Version:   "1.0",
			Service:   &LightDummyAPI{},
			Public:    true,
		}, {
			Namespace: "epv",
			Version:   "1.0",
			Service:   downloader.NewPublicDownloaderAPI(s.protocolManager.downloader, s.eventMux),
			Public:    true,
		}, {
			Namespace: "epv",
			Version:   "1.0",
			Service:   filters.NewPublicFilterAPI(s.ApiBackend, true),
			Public:    true,
		}, {
			Namespace: "net",
			Version:   "1.0",
			Service:   s.netRPCService,
			Public:    true,
		},
	}...)
}

func (s *LightEPVchain) ResetWithGenesisBlock(gb *types.Block) {
	s.blockchain.ResetWithGenesisBlock(gb)
}

func (s *LightEPVchain) BlockChain() *light.LightChain      { return s.blockchain }
func (s *LightEPVchain) TxPool() *light.TxPool              { return s.txPool }
func (s *LightEPVchain) Engine() consensus.Engine           { return s.engine }
func (s *LightEPVchain) LesVersion() int                    { return int(s.protocolManager.SubProtocols[0].Version) }
func (s *LightEPVchain) Downloader() *downloader.Downloader { return s.protocolManager.downloader }
func (s *LightEPVchain) EventMux() *event.TypeMux           { return s.eventMux }

                                                                            
                              
func (s *LightEPVchain) Protocols() []p2p.Protocol {
	return s.protocolManager.SubProtocols
}

                                                                                
                                    
func (s *LightEPVchain) Start(srvr *p2p.Server) error {
	s.startBloomHandlers()
	log.Warn("Light client mode is an experimental feature")
	s.netRPCService = epvapi.NewPublicNetAPI(srvr, s.networkId)
	                                                                      
	protocolVersion := AdvertiseProtocolVersions[0]
	s.serverPool.start(srvr, lesTopic(s.blockchain.Genesis().Hash(), protocolVersion))
	s.protocolManager.Start(s.config.LightPeers)
	return nil
}

                                                                                
                     
func (s *LightEPVchain) Stop() error {
	s.odr.Stop()
	if s.bloomIndexer != nil {
		s.bloomIndexer.Close()
	}
	if s.chtIndexer != nil {
		s.chtIndexer.Close()
	}
	if s.bloomTrieIndexer != nil {
		s.bloomTrieIndexer.Close()
	}
	s.blockchain.Stop()
	s.protocolManager.Stop()
	s.txPool.Stop()

	s.eventMux.Stop()

	time.Sleep(time.Millisecond * 200)
	s.chainDb.Close()
	close(s.shutdownChan)

	return nil
}
