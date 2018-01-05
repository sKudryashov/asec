package main

import (
	"github.com/sKudryashov/asec/fileserver/cmd/servd"
	"github.com/spf13/cobra"
)

var autoTLS bool
var port int

func main() {
	rootCmd := &cobra.Command{
		Use:   "start",
		Short: "start fileserver",
		Run:   cmdHandler,
	}
	rootCmd.PersistentFlags().BoolVar(&autoTLS, "autotls", false, "Launch server with auto TLS")
	rootCmd.PersistentFlags().IntVar(&port, "port", 80, "Launch server with given port")
	rootCmd.Execute()
}

func cmdHandler(cmd *cobra.Command, args []string) {
	servd.StartServer(port, autoTLS)
}
