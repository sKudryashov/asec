package main

import (
	"github.com/sKudryashov/asec/fileserver/servd"
	"github.com/spf13/cobra"
)

var autoTLS bool
var port int
var help bool

// func init() {
// 	autoTLS = flag.Bool("autotls", false, "Launch server with auto TLS *required")
// 	port = flag.Int("port", 80, "Launch server with given port")
// 	help = flag.Bool("help", false, "Print cmd help")

// 	flag.Parse()
// }

func main() {
	rootCmd := &cobra.Command{
		Use:   "filesrv",
		Short: "start fileserver",
		Run:   cmdHandler,
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
