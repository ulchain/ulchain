
package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"unicode"

	cli "gopkg.in/urfave/cli.v1"

	"github.com/epvchain/go-epvchain/command/utils"
	"github.com/epvchain/go-epvchain/panel"
	"github.com/epvchain/go-epvchain/epv"
	"github.com/epvchain/go-epvchain/point"
	"github.com/epvchain/go-epvchain/content"
	whisper "github.com/epvchain/go-epvchain/topic/whisperv5"
	"github.com/naoina/toml"
)

var (
	dumpConfigCommand = cli.Command{
		Action:      utils.MigrateFlags(dumpConfig),
		Name:        "dumpconfig",
		Usage:       "Show configuration values",
		ArgsUsage:   "",
		Flags:       append(append(nodeFlags, rpcFlags...), whisperFlags...),
		Category:    "MISCELLANEOUS COMMANDS",
		Description: `The dumpconfig command shows configuration values.`,
	}

	configFileFlag = cli.StringFlag{
		Name:  "config",
		Usage: "TOML configuration file",
	}
)

var tomlSettings = toml.Config{
	NormFieldName: func(rt reflect.Type, key string) string {
		return key
	},
	FieldToKey: func(rt reflect.Type, field string) string {
		return field
	},
	MissingField: func(rt reflect.Type, field string) error {
		link := ""
		if unicode.IsUpper(rune(rt.Name()[0])) && rt.PkgPath() != "main" {
			link = fmt.Sprintf(", see https://godoc.org/%s#%s for available fields", rt.PkgPath(), rt.Name())
		}
		return fmt.Errorf("field '%s' is not defined in %s%s", field, rt.String(), link)
	},
}

type epvstatsConfig struct {
	URL string `toml:",omitempty"`
}

type gepvConfig struct {
	EPV       epv.Config
	Shh       whisper.Config
	Node      node.Config
	EPVstats  epvstatsConfig
	Dashboard dashboard.Config
}

func loadConfig(file string, cfg *gepvConfig) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	err = tomlSettings.NewDecoder(bufio.NewReader(f)).Decode(cfg)

	if _, ok := err.(*toml.LineError); ok {
		err = errors.New(file + ", " + err.Error())
	}
	return err
}

func defaultNodeConfig() node.Config {
	cfg := node.DefaultConfig
	cfg.Name = clientIdentifier
	cfg.Version = params.VersionWithCommit(gitCommit)
	cfg.HTTPModules = append(cfg.HTTPModules, "epv", "shh")
	cfg.WSModules = append(cfg.WSModules, "epv", "shh")
	cfg.IPCPath = "gepv.ipc"
	return cfg
}

func makeConfigNode(ctx *cli.Context) (*node.Node, gepvConfig) {

	cfg := gepvConfig{
		EPV:       epv.DefaultConfig,
		Shh:       whisper.DefaultConfig,
		Node:      defaultNodeConfig(),
		Dashboard: dashboard.DefaultConfig,
	}

	if file := ctx.GlobalString(configFileFlag.Name); file != "" {
		if err := loadConfig(file, &cfg); err != nil {
			utils.Fatalf("%v", err)
		}
	}

	utils.SetNodeConfig(ctx, &cfg.Node)
	stack, err := node.New(&cfg.Node)
	if err != nil {
		utils.Fatalf("Failed to create the protocol stack: %v", err)
	}
	utils.SetEPVConfig(ctx, stack, &cfg.EPV)
	if ctx.GlobalIsSet(utils.EPVStatsURLFlag.Name) {
		cfg.EPVstats.URL = ctx.GlobalString(utils.EPVStatsURLFlag.Name)
	}

	utils.SetShhConfig(ctx, stack, &cfg.Shh)
	utils.SetDashboardConfig(ctx, &cfg.Dashboard)

	return stack, cfg
}

func enableWhisper(ctx *cli.Context) bool {
	for _, flag := range whisperFlags {
		if ctx.GlobalIsSet(flag.GetName()) {
			return true
		}
	}
	return false
}

func makeFullNode(ctx *cli.Context) *node.Node {
	stack, cfg := makeConfigNode(ctx)

	utils.RegisterEPVService(stack, &cfg.EPV)

	if ctx.GlobalBool(utils.DashboardEnabledFlag.Name) {
		utils.RegisterDashboardService(stack, &cfg.Dashboard, gitCommit)
	}

	shhEnabled := enableWhisper(ctx)
	shhAutoEnabled := !ctx.GlobalIsSet(utils.WhisperEnabledFlag.Name) && ctx.GlobalIsSet(utils.DeveloperFlag.Name)
	if shhEnabled || shhAutoEnabled {
		if ctx.GlobalIsSet(utils.WhisperMaxMessageSizeFlag.Name) {
			cfg.Shh.MaxMessageSize = uint32(ctx.Int(utils.WhisperMaxMessageSizeFlag.Name))
		}
		if ctx.GlobalIsSet(utils.WhisperMinPOWFlag.Name) {
			cfg.Shh.MinimumAcceptedPOW = ctx.Float64(utils.WhisperMinPOWFlag.Name)
		}
		utils.RegisterShhService(stack, &cfg.Shh)
	}

	if cfg.EPVstats.URL != "" {
		utils.RegisterEPVStatsService(stack, cfg.EPVstats.URL)
	}
	return stack
}

func dumpConfig(ctx *cli.Context) error {
	_, cfg := makeConfigNode(ctx)
	comment := ""

	if cfg.EPV.Genesis != nil {
		cfg.EPV.Genesis = nil
		comment += "# Note: this config doesn't contain the genesis block.\n\n"
	}

	out, err := tomlSettings.Marshal(&cfg)
	if err != nil {
		return err
	}
	io.WriteString(os.Stdout, comment)
	os.Stdout.Write(out)
	return nil
}
