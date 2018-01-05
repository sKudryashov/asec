package main

import (
	"log"
	"time"
)

func main() {
	for {
		time.Sleep(5 * time.Second)
		log.Print("server daemon tick")
	}
}
