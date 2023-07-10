package main

import (
	"context"
	"log"
	"math/rand"
	"time"
)

func main() {
	log.Println("START")
	ctx1, cancel1 := context.WithCancel(context.Background())

	ctx2, cancel2 := context.WithCancel(ctx1)

	ctx3, cancel3 := context.WithCancel(context.Background())

	defer cancel1()
	defer cancel2()
	defer cancel3()

	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go func() {
		n := rand.Intn(1000) // n will be between 0 and 10
		log.Println("CTX1 DONE Sleeping ", n, " seconds...")
		time.Sleep(time.Duration(n) * time.Millisecond)
		ctx1.Done()
		log.Println("CTX1 Done")
	}()

	go func() {
		n := rand.Intn(1000) // n will be between 0 and 10
		log.Println("CTX1 Cancel Sleeping ", n, " seconds...")
		time.Sleep(time.Duration(n) * time.Millisecond)
		// cancel1()
		ch1 <- "CTX1 RESULT"
		log.Println("CTX1 Cancel")
	}()

	go func() {
		n := rand.Intn(1000) // n will be between 0 and 10
		log.Println("CTX2 Done Sleeping ", n, " seconds...")
		time.Sleep(time.Duration(n) * time.Millisecond)
		ctx2.Done()
		log.Println("CTX2 Done")
	}()

	go func() {
		n := rand.Intn(1000) // n will be between 0 and 10
		log.Println("CTX2 cancel Sleeping ", n, " seconds...")
		time.Sleep(time.Duration(n) * time.Millisecond)
		ch2 <- "CTX2 RESULT"
		log.Println("CTX2 cancel")
	}()

	go func() {
		n := rand.Intn(1000) // n will be between 0 and 10
		log.Println("CTX3 Done Sleeping ", n, " seconds...")
		time.Sleep(time.Duration(n) * time.Millisecond)
		ctx3.Done()
		log.Println("CTX3 Done")
	}()

	go func() {
		n := rand.Intn(1000) // n will be between 0 and 10
		log.Println("CTX3 cancel Sleeping ", n, " seconds...")
		time.Sleep(time.Duration(n) * time.Millisecond)
		ch3 <- "CTX3 RESULT"
		log.Println("CTX3 cancel")
	}()

	// Wait for the long-running operation to finish or be canceled
	select {
	case res := <-ch1:
		log.Println(res)
	case <-ctx1.Done():
		log.Println("main canceled")
	}

	select {
	case res := <-ch2:
		log.Println(res)
	case <-ctx2.Done():
		log.Println("main canceled")
	}

	select {
	case res := <-ch3:
		log.Println(res)
	case <-ctx3.Done():
		log.Println("main canceled")
	}

	time.Sleep(20 * time.Second)
	log.Println("END")
}
