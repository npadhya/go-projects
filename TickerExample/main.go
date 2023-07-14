package main

import (
	"log"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)

	for {
		log.Println("Starting Ticker")
		time.Sleep(1 * time.Millisecond)
		<-ticker.C
	}
}
