package main

import (
	"fmt"

	"github.com/sKudryashov/asec/fileserver/servd"
	"github.com/spf13/cobra"
)

var autoTLS bool
var port int
var help bool

func main() {
	rootCmd := &cobra.Command{
		Use:                "filesrv",
		Short:              "start fileserver",
		Run:                cmdHandler,
		Args:               argsHandler,
		DisableFlagParsing: true,
	}
	rootCmd.PersistentFlags().BoolVar(&autoTLS, "autotls", false, "Launch server with auto TLS")
	rootCmd.PersistentFlags().IntVar(&port, "port", 80, "Launch server with given port")
	rootCmd.Execute()
}

func cmdHandler(cmd *cobra.Command, args []string) {
	if &port == nil || &autoTLS == nil {
		// flag.PrintDefaults()
	} else {
		servd.StartServer(port, autoTLS)
	}
}

func argsHandler(cmd *cobra.Command, args []string) error {
	if len(args) < 3 {
		return fmt.Errorf("too few args")
	}
	return nil
}
