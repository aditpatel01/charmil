package main

import (
	"log"

	"github.com/aerogear/charmil/examples/plugins/adder"
	"github.com/aerogear/charmil/examples/plugins/date"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "Host",
		Short:        "Host CLI for embedding commands",
		SilenceUsage: true,
	}

	dateCmd, err := date.DateCommand()
	if err != nil {
		log.Fatal(err)
	}

	adderCmd, err := adder.AdderCommand()
	if err != nil {
		log.Fatal(err)
	}

	cmd.AddCommand(dateCmd)
	cmd.AddCommand(adderCmd)

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
