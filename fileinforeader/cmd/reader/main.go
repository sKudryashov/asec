package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sKudryashov/asec/fileinforeader/model"
)

var url string
var port *int
var schema *string
var help *bool

func init() {
	url = "fileserver"
}

func main() {
	port = flag.Int("port", 0, "fetch from server with given port *required")
	schema = flag.String("schema", "http", "schema, http(default) or https used on server")
	help = flag.Bool("help", false, "print help")
	flag.Parse()

	if *help == true {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *port == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	done := make(chan struct{})
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go statListener(sigChan, done)
	log.Println("Listening for events:")
loop:
	for {
		select {
		case <-done:
			break loop
		default:
			time.Sleep(3 * time.Second)
			rsp, err := http.Get(fmt.Sprintf("%s://%s:%d/", *schema, url, *port))
			if err != nil {
				log.Printf("error getting response: %v", err)
			}
			body := readBody(rsp)
			for num, bodyItem := range body {
				log.Printf("record num %d)  name: %s, mode: %s, modeTime: %s, size: %d", num, bodyItem.Name, bodyItem.Mode, bodyItem.ModTime, bodyItem.Size)
			}
			log.Println(".")
			rsp.Body.Close()
		}
	}
}

type stat struct {
	Stat int `json:"rps"`
}

func publishServerStat() {
	log.Println("publishServerStat called")
	rsp, err := http.Get(fmt.Sprintf("%s://%s:%d/stat/", *schema, url, *port))
	if err != nil {
		log.Printf("error acquiring stat %v", err)
	}
	defer rsp.Body.Close()
	st := &stat{}
	if json.NewDecoder(rsp.Body).Decode(st); err != nil {
		log.Printf("error decoding stat body %v", err)
	}
	log.Printf("Server stat (number of records per sec): %d", st.Stat)
}

func statListener(c <-chan os.Signal, done chan<- struct{}) {
	log.Println("listening for ctr+c")
	<-c
	log.Println("ctrl+c pressed")
	publishServerStat()
	done <- struct{}{}
	os.Exit(1)
}

func readBody(rsp *http.Response) []model.FileInfo {
	finfo := make([]model.FileInfo, 0, 15)
	if err := json.NewDecoder(rsp.Body).Decode(&finfo); err != nil {
		log.Printf("error reading body %v", err)
	}
	return finfo
}
