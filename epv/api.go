                                         
                                                
  
                                                                                  
                                                                              
                                                                    
                                      
  
                                                                             
                                                                 
                                                               
                                                      
  
                                                                           
                                                                                  

package epv

import (
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"math/big"
	"os"
	"strings"

	"github.com/epvchain/go-epvchain/public"
	"github.com/epvchain/go-epvchain/public/hexutil"
	"github.com/epvchain/go-epvchain/kernel"
	"github.com/epvchain/go-epvchain/kernel/state"
	"github.com/epvchain/go-epvchain/kernel/types"
	"github.com/epvchain/go-epvchain/book"
	"github.com/epvchain/go-epvchain/work"
	"github.com/epvchain/go-epvchain/content"
	"github.com/epvchain/go-epvchain/process"
	"github.com/epvchain/go-epvchain/remote"
	"github.com/epvchain/go-epvchain/fast"
)

                                                                         
               
type PublicEPVchainAPI struct {
	e *EPVchain
}

                                                                           
func NewPublicEPVchainAPI(e *EPVchain) *PublicEPVchainAPI {
	return &PublicEPVchainAPI{e}
}

                                                              
func (api *PublicEPVchainAPI) EPVCbase() (common.Address, error) {
	return api.e.EPVCbase()
}

                                                                                   
func (api *PublicEPVchainAPI) Coinbase() (common.Address, error) {
	return api.EPVCbase()
}

                                    
func (api *PublicEPVchainAPI) Hashrate() hexutil.Uint64 {
	return hexutil.Uint64(api.e.Miner().HashRate())
}

                                                       
                                                                                                         
type PublicMinerAPI struct {
	e     *EPVchain
	agent *miner.RemoteAgent
}

                                                          
func NewPublicMinerAPI(e *EPVchain) *PublicMinerAPI {
	agent := miner.NewRemoteAgent(e.BlockChain(), e.Engine())
	e.Miner().Register(agent)

	return &PublicMinerAPI{e, agent}
}

                                                                 
func (api *PublicMinerAPI) Mining() bool {
	return api.e.IsMining()
}

                                                                                                                  
                                                                            
func (api *PublicMinerAPI) SubmitWork(nonce types.BlockNonce, solution, digest common.Hash) bool {
	return api.agent.SubmitWork(nonce, digest, solution)
}

                                                                                            
                                                                
                                                         
                                                                                  
func (api *PublicMinerAPI) GetWork() ([3]string, error) {
	if !api.e.IsMining() {
		if err := api.e.StartMining(false); err != nil {
			return [3]string{}, err
		}
	}
	work, err := api.agent.GetWork()
	if err != nil {
		return work, fmt.Errorf("mining not ready: %v", err)
	}
	return work, nil
}

                                                                                                                       
                                                                                                                      
                                
func (api *PublicMinerAPI) SubmitHashrate(hashrate hexutil.Uint64, id common.Hash) bool {
	api.agent.SubmitHashrate(id, uint64(hashrate))
	return true
}

                                                                     
                                                                                                            
type PrivateMinerAPI struct {
	e *EPVchain
}

                                                                                     
func NewPrivateMinerAPI(e *EPVchain) *PrivateMinerAPI {
	return &PrivateMinerAPI{e: e}
}

                                                                                 
                                                                               
                                                                               
                          
func (api *PrivateMinerAPI) Start(threads *int) error {
	                                                           
	if threads == nil {
		threads = new(int)
	} else if *threads == 0 {
		*threads = -1                                 
	}
	type threaded interface {
		SetThreads(threads int)
	}
	if th, ok := api.e.engine.(threaded); ok {
		log.Info("Updated mining threads", "threads", *threads)
		th.SetThreads(*threads)
	}
	                             
	if !api.e.IsMining() {
		                                                            
		api.e.lock.RLock()
		price := api.e.gasPrice
		api.e.lock.RUnlock()

		api.e.txPool.SetGasPrice(price)
		return api.e.StartMining(true)
	}
	return nil
}

                 
func (api *PrivateMinerAPI) Stop() bool {
	type threaded interface {
		SetThreads(threads int)
	}
	if th, ok := api.e.engine.(threaded); ok {
		th.SetThreads(-1)
	}
	api.e.StopMining()
	return true
}

                                                                                      
func (api *PrivateMinerAPI) SetExtra(extra string) (bool, error) {
	if err := api.e.Miner().SetExtra([]byte(extra)); err != nil {
		return false, err
	}
	return true, nil
}

                                                                 
func (api *PrivateMinerAPI) SetGasPrice(gasPrice hexutil.Big) bool {
	api.e.lock.Lock()
	api.e.gasPrice = (*big.Int)(&gasPrice)
	api.e.lock.Unlock()

	api.e.txPool.SetGasPrice((*big.Int)(&gasPrice))
	return true
}

                                             
func (api *PrivateMinerAPI) SetEPVCbase(epvcbase common.Address) bool {
	api.e.SetEPVCbase(epvcbase)
	return true
}

                                                         
func (api *PrivateMinerAPI) GetHashrate() uint64 {
	return uint64(api.e.miner.HashRate())
}

                                                                       
                                           
type PrivateAdminAPI struct {
	epv *EPVchain
}

                                                                            
                                         
func NewPrivateAdminAPI(epv *EPVchain) *PrivateAdminAPI {
	return &PrivateAdminAPI{epv: epv}
}

                                                                
func (api *PrivateAdminAPI) ExportChain(file string) (bool, error) {
	                                                  
	out, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return false, err
	}
	defer out.Close()

	var writer io.Writer = out
	if strings.HasSuffix(file, ".gz") {
		writer = gzip.NewWriter(writer)
		defer writer.(*gzip.Writer).Close()
	}

	                        
	if err := api.epv.BlockChain().Export(writer); err != nil {
		return false, err
	}
	return true, nil
}

func hasAllBlocks(chain *core.BlockChain, bs []*types.Block) bool {
	for _, b := range bs {
		if !chain.HasBlock(b.Hash(), b.NumberU64()) {
			return false
		}
	}

	return true
}

                                                      
func (api *PrivateAdminAPI) ImportChain(file string) (bool, error) {
	                                              
	in, err := os.Open(file)
	if err != nil {
		return false, err
	}
	defer in.Close()

	var reader io.Reader = in
	if strings.HasSuffix(file, ".gz") {
		if reader, err = gzip.NewReader(reader); err != nil {
			return false, err
		}
	}

	                                                  
	stream := rlp.NewStream(reader, 0)

	blocks, index := make([]*types.Block, 0, 2500), 0
	for batch := 0; ; batch++ {
		                                             
		for len(blocks) < cap(blocks) {
			block := new(types.Block)
			if err := stream.Decode(block); err == io.EOF {
				break
			} else if err != nil {
				return false, fmt.Errorf("block %d: failed to parse: %v", index, err)
			}
			blocks = append(blocks, block)
			index++
		}
		if len(blocks) == 0 {
			break
		}

		if hasAllBlocks(api.epv.BlockChain(), blocks) {
			blocks = blocks[:0]
			continue
		}
		                                        
		if _, err := api.epv.BlockChain().InsertChain(blocks); err != nil {
			return false, fmt.Errorf("batch %d: failed to insert: %v", batch, err)
		}
		blocks = blocks[:0]
	}
	return true, nil
}

                                                                      
                                      
type PublicDebugAPI struct {
	epv *EPVchain
}

                                                                    
                                                        
func NewPublicDebugAPI(epv *EPVchain) *PublicDebugAPI {
	return &PublicDebugAPI{epv: epv}
}

                                                                         
func (api *PublicDebugAPI) DumpBlock(blockNr rpc.BlockNumber) (state.Dump, error) {
	if blockNr == rpc.PendingBlockNumber {
		                                                         
		                                                           
		                                 
		_, stateDb := api.epv.miner.Pending()
		return stateDb.RawDump(), nil
	}
	var block *types.Block
	if blockNr == rpc.LatestBlockNumber {
		block = api.epv.blockchain.CurrentBlock()
	} else {
		block = api.epv.blockchain.GetBlockByNumber(uint64(blockNr))
	}
	if block == nil {
		return state.Dump{}, fmt.Errorf("block #%d not found", blockNr)
	}
	stateDb, err := api.epv.BlockChain().StateAt(block.Root())
	if err != nil {
		return state.Dump{}, err
	}
	return stateDb.RawDump(), nil
}

                                                                            
                                  
type PrivateDebugAPI struct {
	config *params.ChainConfig
	epv    *EPVchain
}

                                                                            
                                                 
func NewPrivateDebugAPI(config *params.ChainConfig, epv *EPVchain) *PrivateDebugAPI {
	return &PrivateDebugAPI{config: config, epv: epv}
}

                                                                                        
func (api *PrivateDebugAPI) Preimage(ctx context.Context, hash common.Hash) (hexutil.Bytes, error) {
	db := core.PreimageTable(api.epv.ChainDb())
	return db.Get(hash.Bytes())
}

                                                                                               
                                                  
func (api *PrivateDebugAPI) GetBadBlocks(ctx context.Context) ([]core.BadBlockArgs, error) {
	return api.epv.BlockChain().BadBlocks()
}

                                                                       
type StorageRangeResult struct {
	Storage storageMap   `json:"storage"`
	NextKey *common.Hash `json:"nextKey"`                                                     
}

type storageMap map[common.Hash]storageEntry

type storageEntry struct {
	Key   *common.Hash `json:"key"`
	Value common.Hash  `json:"value"`
}

                                                                                      
func (api *PrivateDebugAPI) StorageRangeAt(ctx context.Context, blockHash common.Hash, txIndex int, contractAddress common.Address, keyStart hexutil.Bytes, maxResult int) (StorageRangeResult, error) {
	_, _, statedb, err := api.computeTxEnv(blockHash, txIndex, 0)
	if err != nil {
		return StorageRangeResult{}, err
	}
	st := statedb.StorageTrie(contractAddress)
	if st == nil {
		return StorageRangeResult{}, fmt.Errorf("account %x doesn't exist", contractAddress)
	}
	return storageRangeAt(st, keyStart, maxResult)
}

func storageRangeAt(st state.Trie, start []byte, maxResult int) (StorageRangeResult, error) {
	it := trie.NewIterator(st.NodeIterator(start))
	result := StorageRangeResult{Storage: storageMap{}}
	for i := 0; i < maxResult && it.Next(); i++ {
		_, content, _, err := rlp.Split(it.Value)
		if err != nil {
			return StorageRangeResult{}, err
		}
		e := storageEntry{Value: common.BytesToHash(content)}
		if preimage := st.GetKey(it.Key); preimage != nil {
			preimage := common.BytesToHash(preimage)
			e.Key = &preimage
		}
		result.Storage[common.BytesToHash(it.Key)] = e
	}
	                                                          
	if it.Next() {
		next := common.BytesToHash(it.Key)
		result.NextKey = &next
	}
	return result, nil
}

                                                                                
                                                                               
                              
  
                                                                                    
func (api *PrivateDebugAPI) GetModifiedAccountsByNumber(startNum uint64, endNum *uint64) ([]common.Address, error) {
	var startBlock, endBlock *types.Block

	startBlock = api.epv.blockchain.GetBlockByNumber(startNum)
	if startBlock == nil {
		return nil, fmt.Errorf("start block %x not found", startNum)
	}

	if endNum == nil {
		endBlock = startBlock
		startBlock = api.epv.blockchain.GetBlockByHash(startBlock.ParentHash())
		if startBlock == nil {
			return nil, fmt.Errorf("block %x has no parent", endBlock.Number())
		}
	} else {
		endBlock = api.epv.blockchain.GetBlockByNumber(*endNum)
		if endBlock == nil {
			return nil, fmt.Errorf("end block %d not found", *endNum)
		}
	}
	return api.getModifiedAccounts(startBlock, endBlock)
}

                                                                               
                                                                               
                              
  
                                                                                    
func (api *PrivateDebugAPI) GetModifiedAccountsByHash(startHash common.Hash, endHash *common.Hash) ([]common.Address, error) {
	var startBlock, endBlock *types.Block
	startBlock = api.epv.blockchain.GetBlockByHash(startHash)
	if startBlock == nil {
		return nil, fmt.Errorf("start block %x not found", startHash)
	}

	if endHash == nil {
		endBlock = startBlock
		startBlock = api.epv.blockchain.GetBlockByHash(startBlock.ParentHash())
		if startBlock == nil {
			return nil, fmt.Errorf("block %x has no parent", endBlock.Number())
		}
	} else {
		endBlock = api.epv.blockchain.GetBlockByHash(*endHash)
		if endBlock == nil {
			return nil, fmt.Errorf("end block %x not found", *endHash)
		}
	}
	return api.getModifiedAccounts(startBlock, endBlock)
}

func (api *PrivateDebugAPI) getModifiedAccounts(startBlock, endBlock *types.Block) ([]common.Address, error) {
	if startBlock.Number().Uint64() >= endBlock.Number().Uint64() {
		return nil, fmt.Errorf("start block height (%d) must be less than end block height (%d)", startBlock.Number().Uint64(), endBlock.Number().Uint64())
	}

	oldTrie, err := trie.NewSecure(startBlock.Root(), trie.NewDatabase(api.epv.chainDb), 0)
	if err != nil {
		return nil, err
	}
	newTrie, err := trie.NewSecure(endBlock.Root(), trie.NewDatabase(api.epv.chainDb), 0)
	if err != nil {
		return nil, err
	}

	diff, _ := trie.NewDifferenceIterator(oldTrie.NodeIterator([]byte{}), newTrie.NodeIterator([]byte{}))
	iter := trie.NewIterator(diff)

	var dirty []common.Address
	for iter.Next() {
		key := newTrie.GetKey(iter.Key)
		if key == nil {
			return nil, fmt.Errorf("no preimage found for hash %x", iter.Key)
		}
		dirty = append(dirty, common.BytesToAddress(key))
	}
	return dirty, nil
}
