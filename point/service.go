                                         
                                                
  
                                                                                  
                                                                              
                                                                    
                                      
  
                                                                             
                                                                 
                                                               
                                                      
  
                                                                           
                                                                                  

package node

import (
	"reflect"

	"github.com/epvchain/go-epvchain/act"
	"github.com/epvchain/go-epvchain/data"
	"github.com/epvchain/go-epvchain/notice"
	"github.com/epvchain/go-epvchain/peer"
	"github.com/epvchain/go-epvchain/remote"
)

                                                                               
                                                                                
                                                                    
type ServiceContext struct {
	config         *Config
	services       map[reflect.Type]Service                                             
	EventMux       *event.TypeMux                                                                
	AccountManager *accounts.Manager                                               
}

                                                                              
                                                                             
                                                           
func (ctx *ServiceContext) OpenDatabase(name string, cache int, handles int) (epvdb.Database, error) {
	if ctx.config.DataDir == "" {
		return epvdb.NewMemDatabase()
	}
	db, err := epvdb.NewLDBDatabase(ctx.config.resolvePath(name), cache, handles)
	if err != nil {
		return nil, err
	}
	return db, nil
}

                                                                                
                                                                                   
                                                                      
func (ctx *ServiceContext) ResolvePath(path string) string {
	return ctx.config.resolvePath(path)
}

                                                                               
func (ctx *ServiceContext) Service(service interface{}) error {
	element := reflect.ValueOf(service).Elem()
	if running, ok := ctx.services[element.Type()]; ok {
		element.Set(reflect.ValueOf(running))
		return nil
	}
	return ErrServiceUnknown
}

                                                                                
                                        
type ServiceConstructor func(ctx *ServiceContext) (Service, error)

                                                                        
  
         
  
                                                                                        
                                                                                      
                
  
                                                                             
                                   
type Service interface {
	                                                                     
	Protocols() []p2p.Protocol

	                                                                  
	APIs() []rpc.API

	                                                                              
	                                                                              
	Start(server *p2p.Server) error

	                                                                               
	                      
	Stop() error
}
