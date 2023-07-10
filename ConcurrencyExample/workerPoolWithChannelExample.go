package main

import "log"

func doWork() {
	jobs := make(chan int, 20)
	results := make(chan int, 20)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 20; i++ {
		jobs <- i
	}

	close(jobs)

	for j := 0; j < 20; j++ {
		log.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}
}
