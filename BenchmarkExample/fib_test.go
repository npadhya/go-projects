package main

import (
	"testing"
)

func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(40)
	}
}

func fib(n int) int {
	if n <= 1 {

		return n
	}
	return fib(n-1) + fib(n-2)
}
