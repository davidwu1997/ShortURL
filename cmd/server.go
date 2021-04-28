package cmd

import (
	"fmt"
	"os"
	"shortURL/pkg/service/shorturl"

	cobra "github.com/spf13/cobra"
)

// ServerCmd http server
var ServerCmd = &cobra.Command{
	Run:           runServerCmd,
	Use:           "server",
	Short:         "Start ShortURL server",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func runServerCmd(cmd *cobra.Command, args []string) {
	app, err := shorturl.Initialize()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	err = app.Start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
