                                         
                                                
  
                                                                                  
                                                                              
                                                                    
                                      
  
                                                                             
                                                                 
                                                               
                                                      
  
                                                                           
                                                                                  

package core

import (
	"github.com/epvchain/go-epvchain/public"
	"github.com/epvchain/go-epvchain/kernel/types"
)

                                                                       
type TxPreEvent struct{ Tx *types.Transaction }

                                                                      
type PendingLogsEvent struct {
	Logs []*types.Log
}

                                                                                
type PendingStateEvent struct{}

                                                               
type NewMinedBlockEvent struct{ Block *types.Block }

                                                         
type RemovedTransactionEvent struct{ Txs types.Transactions }

                                                  
type RemovedLogsEvent struct{ Logs []*types.Log }

type ChainEvent struct {
	Block *types.Block
	Hash  common.Hash
	Logs  []*types.Log
}

type ChainSideEvent struct {
	Block *types.Block
}

type ChainHeadEvent struct{ Block *types.Block }
