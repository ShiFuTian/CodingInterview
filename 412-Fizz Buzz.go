package leetcode

import (
	"strconv"
)

func fizzBuzz(n int) []string {
	var ret = make([]string, n)
	for i := 1; i <= n; i++ {
		ret[i-1] = strconv.Itoa(i)
	}
	for i := 3; i <= n; i += 3 {
		ret[i-1] = "Fizz"
	}
	for i := 5; i <= n; i += 5 {
		ret[i-1] = "Buzz"
	}
	for i := 15; i <= n; i += 15 {
		ret[i-1] = "FizzBuzz"
	}
	return ret
}

func fizzBuzz2(n int) []string {
	var ret []string
	for i := 1; i <= n; i++ {
		ret = append(ret, strconv.FormatInt(int64(i), 10))
	}
	for i := 3; i <= n; i += 3 {
		ret[i-1] = "Fizz"
	}
	for i := 5; i <= n; i += 5 {
		ret[i-1] = "Buzz"
	}
	for i := 15; i <= n; i += 15 {
		ret[i-1] = "FizzBuzz"
	}
	return ret
}
