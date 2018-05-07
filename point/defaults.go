
package node

import (
	"os"
	"os/user"
	"path/filepath"
	"runtime"

	"github.com/epvchain/go-epvchain/peer"
	"github.com/epvchain/go-epvchain/peer/nat"
)

const (
	DefaultHTTPHost = "localhost" 
	DefaultHTTPPort = 7545        
	DefaultWSHost   = "localhost" 
	DefaultWSPort   = 7546        
)

var DefaultConfig = Config{
	DataDir:     DefaultDataDir(),
	HTTPPort:    DefaultHTTPPort,
	HTTPModules: []string{"net", "web3"},
	WSPort:      DefaultWSPort,
	WSModules:   []string{"net", "web3"},
	P2P: p2p.Config{
		ListenAddr: ":50303",
		MaxPeers:   25,
		NAT:        nat.Any(),
	},
}

func DefaultDataDir() string {

	home := homeDir()
	if home != "" {
		if runtime.GOOS == "darwin" {
			return filepath.Join(home, "Library", "EPVchain")
		} else if runtime.GOOS == "windows" {
			return filepath.Join(home, "AppData", "Roaming", "EPVchain")
		} else {
			return filepath.Join(home, ".epvchain")
		}
	}

	return ""
}

func homeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	if usr, err := user.Current(); err == nil {
		return usr.HomeDir
	}
	return ""
}
