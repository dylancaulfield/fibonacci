package main

import "testing"

func BenchmarkFibRecursive(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		fibRecursive(30)
	}
}

func BenchmarkFibRecursiveMemo(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		fibRecursiveMemo(30)
	}
}

func BenchmarkFibInOrder(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		fibInOrder(30)
	}
}