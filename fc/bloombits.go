                                         
                                                
  
                                                                                  
                                                                              
                                                                    
                                      
  
                                                                             
                                                                 
                                                               
                                                      
  
                                                                           
                                                                                  

package les

import (
	"time"

	"github.com/epvchain/go-epvchain/public/bitutil"
	"github.com/epvchain/go-epvchain/simple"
)

const (
	                                                                               
	                                                                 
	bloomServiceThreads = 16

	                                                                            
	                                                           
	bloomFilterThreads = 3

	                                                                               
	                     
	bloomRetrievalBatch = 16

	                                                                               
	                                                               
	bloomRetrievalWait = time.Microsecond * 100
)

                                                                               
                                                                               
func (epv *LightEPVchain) startBloomHandlers() {
	for i := 0; i < bloomServiceThreads; i++ {
		go func() {
			for {
				select {
				case <-epv.shutdownChan:
					return

				case request := <-epv.bloomRequests:
					task := <-request
					task.Bitsets = make([][]byte, len(task.Sections))
					compVectors, err := light.GetBloomBits(task.Context, epv.odr, task.Bit, task.Sections)
					if err == nil {
						for i := range task.Sections {
							if blob, err := bitutil.DecompressBytes(compVectors[i], int(light.BloomTrieFrequency/8)); err == nil {
								task.Bitsets[i] = blob
							} else {
								task.Error = err
							}
						}
					} else {
						task.Error = err
					}
					request <- task
				}
			}
		}()
	}
}

const (
	                                                                               
	                                                                 
	bloomConfirms = 256

	                                                                               
	                                                                        
	bloomThrottling = 100 * time.Millisecond
)
