package leetcode

import "testing"

func Benchmark_fizzBuzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fizzBuzz(50)
	}
}

func Benchmark_fizzBuzz2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fizzBuzz(50)
	}
}
