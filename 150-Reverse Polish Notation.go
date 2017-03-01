package leetcode

import (
	"log"
	"strconv"
	"strings"
)

// Problem 150
// ["2", "1", "+", "3", "*"] -> ((2 + 1)*3) -> 9
// ["4", "13", "5", "/", "+"] -> (4 + (13 / 5)) -> 6

func evalRPN(tokens []string) int {
	const ops = "+-*/" // operators
	var stack []int    // use stack

	for _, token := range tokens {

		if strings.Contains(ops, token) { // token is operator
			slen := len(stack)
			if slen < 2 { // 1 operator with 0 or 1 operand, invalid
				log.Fatal("invalid input")
			}

			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:slen-2] // pop a, b from stack

			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				stack = append(stack, a/b)
			}
			continue
		}

		value, err := strconv.Atoi(token) // parse the operand
		if err != nil {
			log.Fatalf("invalid input, can not parse %s to int", token)
		}

		stack = append(stack, value) // push value to stack
	}

	if len(stack) != 1 { // check whether stack only contains the result or not
		log.Fatalf("invalid input")
	}

	ret := stack[0]
	return ret
}
