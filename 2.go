// 2. Evaluate Reverse Polish Notation
// Evaluate the value of an arithmetic expression in Reverse Polish Notation.
// Valid operators are +, -, *, /. Each operand may be an integer or another expression.
// For example:
//   ["2", "1", "+", "3", "*"] -> ((2 + 1) * 3) -> 9
//   ["4", "13", "5", "/", "+"] -> (4 + (13 / 5)) -> 6
package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Stack interface {
	Push(v interface{})
	Pop() interface{}
}

type IntStack struct {
	slice []int
}

func (i *IntStack) Push(v interface{}) {
	i.slice = append(i.slice, v.(int))
}

func (i *IntStack) Pop() interface{} {
	top := i.slice[len(i.slice)-1]
	i.slice = i.slice[:len(i.slice)-1]
	return top
}

func (i IntStack) String() string {
	var buf bytes.Buffer
	buf.WriteString("[")
	for k, e := range i.slice {
		a := strconv.Itoa(e) // Here cause special
		buf.WriteString(a)
		if k != len(i.slice)-1 {
			buf.WriteString(" ")
		}
	}
	buf.WriteString("]")
	return buf.String()
}

func EvalRPN(tokens []string) (int, error) {
	op := "+-*/"

	var s Stack = &IntStack{make([]int, 0)}

	for _, v := range tokens {
		if !strings.Contains(op, v) {
			i, err := strconv.Atoi(v)
			if err != nil {
				return 0, fmt.Errorf("EvalRPN(%#v), %v\n", tokens, err)
			}
			s.Push(i) // Push int
		} else {
			r := s.Pop().(int) // Pop int
			l := s.Pop().(int)
			switch v {
			case "+":
				s.Push(l + r) // Push int
			case "-":
				s.Push(l - r)
			case "*":
				s.Push(l * r)
			case "/":
				s.Push(l / r)
			}
		}
		fmt.Println(s)
	}
	return s.Pop().(int), nil
}

func main() {
	tokens := []string{"2", "1", "+", "3", "*"}
	v, err := EvalRPN(tokens)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value of", tokens, "is :", v)
	return
}
