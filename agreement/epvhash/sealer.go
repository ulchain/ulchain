                                         
                                                
  
                                                                                  
                                                                              
                                                                    
                                      
  
                                                                             
                                                                 
                                                               
                                                      
  
                                                                           
                                                                                  

package epvhash

import (
	crand "crypto/rand"
	"math"
	"math/big"
	"math/rand"
	"runtime"
	"sync"

	"github.com/epvchain/go-epvchain/public"
	"github.com/epvchain/go-epvchain/agreement"
	"github.com/epvchain/go-epvchain/kernel/types"
	"github.com/epvchain/go-epvchain/book"
)

                                                                              
                                       
func (epvhash *EPVhash) Seal(chain consensus.ChainReader, block *types.Block, stop <-chan struct{}) (*types.Block, error) {
	                                                                   
	if epvhash.config.PowMode == ModeFake || epvhash.config.PowMode == ModeFullFake {
		header := block.Header()
		header.Nonce, header.MixDigest = types.BlockNonce{}, common.Hash{}
		return block.WithSeal(header), nil
	}
	                                                        
	if epvhash.shared != nil {
		return epvhash.shared.Seal(chain, block, stop)
	}
	                                                             
	abort := make(chan struct{})
	found := make(chan *types.Block)

	epvhash.lock.Lock()
	threads := epvhash.threads
	if epvhash.rand == nil {
		seed, err := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
		if err != nil {
			epvhash.lock.Unlock()
			return nil, err
		}
		epvhash.rand = rand.New(rand.NewSource(seed.Int64()))
	}
	epvhash.lock.Unlock()
	if threads == 0 {
		threads = runtime.NumCPU()
	}
	if threads < 0 {
		threads = 0                                                                         
	}
	var pend sync.WaitGroup
	for i := 0; i < threads; i++ {
		pend.Add(1)
		go func(id int, nonce uint64) {
			defer pend.Done()
			epvhash.mine(block, id, nonce, abort, found)
		}(i, uint64(epvhash.rand.Int63()))
	}
	                                                       
	var result *types.Block
	select {
	case <-stop:
		                                        
		close(abort)
	case result = <-found:
		                                                     
		close(abort)
	case <-epvhash.update:
		                                                    
		close(abort)
		pend.Wait()
		return epvhash.Seal(chain, block, stop)
	}
	                                                        
	pend.Wait()
	return result, nil
}

                                                                                 
                                                       
func (epvhash *EPVhash) mine(block *types.Block, id int, seed uint64, abort chan struct{}, found chan *types.Block) {
	                                    
	var (
		header  = block.Header()
		hash    = header.HashNoNonce().Bytes()
		target  = new(big.Int).Div(maxUint256, header.Difficulty)
		number  = header.Number.Uint64()
		dataset = epvhash.dataset(number)
	)
	                                                                   
	var (
		attempts = int64(0)
		nonce    = seed
	)
	logger := log.New("miner", id)
	logger.Trace("Started epvhash search for new nonces", "seed", seed)
search:
	for {
		select {
		case <-abort:
			                                            
			logger.Trace("EPVhash nonce search aborted", "attempts", nonce-seed)
			epvhash.hashrate.Mark(attempts)
			break search

		default:
			                                                                                     
			attempts++
			if (attempts % (1 << 15)) == 0 {
				epvhash.hashrate.Mark(attempts)
				attempts = 0
			}
			                                      
			digest, result := hashimotoFull(dataset.dataset, hash, nonce)
			if new(big.Int).SetBytes(result).Cmp(target) <= 0 {
				                                                   
				header = types.CopyHeader(header)
				header.Nonce = types.EncodeNonce(nonce)
				header.MixDigest = common.BytesToHash(digest)

				                                            
				select {
				case found <- block.WithSeal(header):
					logger.Trace("EPVhash nonce found and reported", "attempts", nonce-seed, "nonce", nonce)
				case <-abort:
					logger.Trace("EPVhash nonce found but discarded", "attempts", nonce-seed, "nonce", nonce)
				}
				break search
			}
			nonce++
		}
	}
	                                                                           
	                                                        
	runtime.KeepAlive(dataset)
}
