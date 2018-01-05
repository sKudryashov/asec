package main

import (
	"flag"
	"os"

	"github.com/sKudryashov/asec/fileserver/cmd/servcli/servd"
)

// var autoTLS bool
// var port int

// type portFlag struct {
// 	set   bool
// 	value string
// }

// func (pf *portFlag) Set(port string) error {
// 	pf.set = true
// 	pf.value = port
// 	return nil
// }

// func (pf *portFlag) String() string {
// 	return pf.value
// }

// func (pf *portFlag) Value() string {
// 	return pf.value
// }

// var portF portFlag

var autoTLS *bool
var port *int
var help *bool

func main() {
	port = flag.Int("port", 0, "launch server with given port *required")
	help = flag.Bool("help", false, "print help")
	autoTLS = flag.Bool("autotls", false, "launch server with auto TLS *required")
	flag.Parse()

	if *help == true {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *port == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	servd.StartServer(*port, *autoTLS)
}
