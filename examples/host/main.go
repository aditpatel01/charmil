package main

import (
	"log"

	"github.com/aerogear/charmil/core/config"
	"github.com/aerogear/charmil/examples/plugins/adder"
	"github.com/aerogear/charmil/examples/plugins/date"
	"github.com/spf13/cobra"
)

var (
	// Stores an instance of the charmil config handler
	h *config.Handler

	// Stores the local config file settings
	f = config.File{
		Name: "config",
		Type: "yaml",
		Path: "./examples/host",
	}

	// Root command of the host CLI
	cmd = &cobra.Command{
		Use:          "Host",
		Short:        "Host CLI for embedding commands",
		SilenceUsage: true,
	}
)

func init() {
	// Assigns a new instance of the charmil config handler
	h = config.New()

	// Links the handler instance to a local config file
	h.InitFile(f)

	// Loads config values from the local config file
	err := h.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Sets a dummy value into config
	h.SetValue("key4", "val4")

	// Stores the root command of the `date` plugin
	dateCmd, err := date.DateCommand()
	if err != nil {
		log.Fatal(err)
	}

	// Stores the root command and the config map of the `adder` plugin
	adderCmd, adderCfg, err := adder.AdderCommand()
	if err != nil {
		log.Fatal(err)
	}

	// Adds the root command of plugins as child commands
	cmd.AddCommand(dateCmd)
	cmd.AddCommand(adderCmd)

	// Maps the plugin name to its imported config map
	h.SetPluginCfg("adder", adderCfg)

	// Stores config of every imported plugin into the current config
	h.MergePluginCfg()

	// Writes the current config into the local config file
	err = h.Save()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
