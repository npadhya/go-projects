package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	doWork()
	go WaitGroupExample("sheep")
	go WaitGroupExample("goat")
	time.Sleep(time.Second * 5)

}

func count(name string) {
	for i := 1; i < 11; i++ {
		log.Println(i, name)
		// time.Sleep(time.Second)
	}
}

func WaitGroupExample(name string) {
	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		count(name + "1")
		wg.Done()
		count(name + "2")
		wg.Done()
		count(name + "3")
		wg.Done()
		count(name + "4")
		wg.Done()
	}()

	log.Println("Starting")

	wg.Wait()
}
