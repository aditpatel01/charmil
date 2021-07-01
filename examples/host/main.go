package main

import (
	"log"

	"github.com/aerogear/charmil/core/config"
	"github.com/aerogear/charmil/examples/plugins/adder"
	"github.com/aerogear/charmil/examples/plugins/date"
	"github.com/spf13/cobra"
)

var (
	h *config.Handler

	f = config.File{
		Name: "config",
		Type: "yaml",
		Path: "./examples/host",
	}

	cmd = &cobra.Command{
		Use:          "Host",
		Short:        "Host CLI for embedding commands",
		SilenceUsage: true,
	}
)

func init() {
	h = config.New()

	h.InitFile(f)

	err := h.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	h.SetValue("key4", "val4")

	dateCmd, err := date.DateCommand()
	if err != nil {
		log.Fatal(err)
	}

	adderCmd, adderCfg, err := adder.AdderCommand()
	if err != nil {
		log.Fatal(err)
	}

	cmd.AddCommand(dateCmd)
	cmd.AddCommand(adderCmd)

	h.SetPluginCfg("adder", adderCfg)

	h.MergePluginCfg()

	err = h.Save()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
