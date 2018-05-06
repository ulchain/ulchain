                                         
                                                
  
                                                                                  
                                                                              
                                                                    
                                      
  
                                                                             
                                                                 
                                                               
                                                      
  
                                                                           
                                                                                  

                                                 

package fetcher

import (
	"github.com/epvchain/go-epvchain/disk"
)

var (
	propAnnounceInMeter   = metrics.NewMeter("epv/fetcher/prop/announces/in")
	propAnnounceOutTimer  = metrics.NewTimer("epv/fetcher/prop/announces/out")
	propAnnounceDropMeter = metrics.NewMeter("epv/fetcher/prop/announces/drop")
	propAnnounceDOSMeter  = metrics.NewMeter("epv/fetcher/prop/announces/dos")

	propBroadcastInMeter   = metrics.NewMeter("epv/fetcher/prop/broadcasts/in")
	propBroadcastOutTimer  = metrics.NewTimer("epv/fetcher/prop/broadcasts/out")
	propBroadcastDropMeter = metrics.NewMeter("epv/fetcher/prop/broadcasts/drop")
	propBroadcastDOSMeter  = metrics.NewMeter("epv/fetcher/prop/broadcasts/dos")

	headerFetchMeter = metrics.NewMeter("epv/fetcher/fetch/headers")
	bodyFetchMeter   = metrics.NewMeter("epv/fetcher/fetch/bodies")

	headerFilterInMeter  = metrics.NewMeter("epv/fetcher/filter/headers/in")
	headerFilterOutMeter = metrics.NewMeter("epv/fetcher/filter/headers/out")
	bodyFilterInMeter    = metrics.NewMeter("epv/fetcher/filter/bodies/in")
	bodyFilterOutMeter   = metrics.NewMeter("epv/fetcher/filter/bodies/out")
)
