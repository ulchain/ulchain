                                         
                                                
  
                                                                                  
                                                                              
                                                                    
                                      
  
                                                                             
                                                                 
                                                               
                                                      
  
                                                                           
                                                                                  

package epv

import (
	"math/big"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"time"

	"github.com/epvchain/go-epvchain/public"
	"github.com/epvchain/go-epvchain/public/hexutil"
	"github.com/epvchain/go-epvchain/agreement/epvhash"
	"github.com/epvchain/go-epvchain/kernel"
	"github.com/epvchain/go-epvchain/epv/downloader"
	"github.com/epvchain/go-epvchain/epv/gasprice"
	"github.com/epvchain/go-epvchain/content"
)

                                                                            
var DefaultConfig = Config{
	SyncMode: downloader.FastSync,
	EPVhash: epvhash.Config{
		CacheDir:       "epvhash",
		CachesInMem:    2,
		CachesOnDisk:   3,
		DatasetsInMem:  1,
		DatasetsOnDisk: 2,
	},
	NetworkId:     1,
	LightPeers:    100,
	DatabaseCache: 768,
	TrieCache:     256,
	TrieTimeout:   5 * time.Minute,
	GasPrice:      big.NewInt(18 * params.Shannon),

	TxPool: core.DefaultTxPoolConfig,
	GPO: gasprice.Config{
		Blocks:     20,
		Percentile: 60,
	},
}

func init() {
	home := os.Getenv("HOME")
	if home == "" {
		if user, err := user.Current(); err == nil {
			home = user.HomeDir
		}
	}
	if runtime.GOOS == "windows" {
		DefaultConfig.EPVhash.DatasetDir = filepath.Join(home, "AppData", "EPVhash")
	} else {
		DefaultConfig.EPVhash.DatasetDir = filepath.Join(home, ".epvhash")
	}
}

                                                                                                     

type Config struct {
	                                                                 
	                                               
	Genesis *core.Genesis `toml:",omitempty"`

	                   
	NetworkId uint64                                                       
	SyncMode  downloader.SyncMode
	NoPruning bool

	                       
	LightServ  int `toml:",omitempty"`                                                               
	LightPeers int `toml:",omitempty"`                                      

	                   
	SkipBcVersionCheck bool `toml:"-"`
	DatabaseHandles    int  `toml:"-"`
	DatabaseCache      int
	TrieCache          int
	TrieTimeout        time.Duration

	                         
	EPVCbase    common.Address `toml:",omitempty"`
	MinerThreads int            `toml:",omitempty"`
	ExtraData    []byte         `toml:",omitempty"`
	GasPrice     *big.Int

	                  
	EPVhash epvhash.Config

	                           
	TxPool core.TxPoolConfig

	                           
	GPO gasprice.Config

	                                               
	EnablePreimageRecording bool

	                        
	DocRoot string `toml:"-"`
}

type configMarshaling struct {
	ExtraData hexutil.Bytes
}
