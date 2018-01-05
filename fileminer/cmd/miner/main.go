package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
)

const (
	traversalFlowsAmount = 10
	filesAmount          = 1000
	codeRandJSON         = 0
	codeRandXML          = 1
)

var path *string
var schema *string
var port *int
var help *bool
var done = make(chan struct{})
var semaphore = make(chan struct{}, traversalFlowsAmount)

func main() {
	port = flag.Int("port", 0, "launch server with given port *required")
	schema = flag.String("schema", "http", "schema, http(default) or https used on server")
	path = flag.String("path", "", "path to start from *required")
	help = flag.Bool("help", false, "print help")
	flag.Parse()

	if *help == true {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if len(*path) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *port == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go cancelListener(sigChan, done)
	cmdHandler()
}

func cancelListener(chOS <-chan os.Signal, chDone chan<- struct{}) {
	<-chOS
	done <- struct{}{}
	os.Exit(1)
}

func cmdHandler() {
	roots := []string{*path}
	traverseDir(roots)
}
