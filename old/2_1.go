// 2. Evaluate Reverse Polish Notation
// Evaluate the value of an arithmetic expression in Reverse Polish Notation.
// Valid operators are +, -, *, /. Each operand may be an integer or another expression.
// For example:
//   ["2", "1", "+", "3", "*"] -> ((2 + 1) * 3) -> 9
//   ["4", "13", "5", "/", "+"] -> (4 + (13 / 5)) -> 6
package main

import (
	"bytes"
	"container/list"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Stack struct {
	*list.List
}

func (s *Stack) Push(v interface{}) {
	s.PushBack(v)
}

func (s *Stack) Pop() interface{} {
	return s.Remove(s.Back())
}

func (s Stack) String() string {
	var buf bytes.Buffer
	buf.WriteString("[")
	for e := s.Front(); e != nil; e = e.Next() {
		a := strconv.Itoa(e.Value.(int))
		buf.WriteString(a)
		if e != s.Back() {
			buf.WriteString(" ")
		}
	}
	buf.WriteString("]")
	return buf.String()
}

func main() {
	tokens := []string{"2.1", "1", "+", "3", "*"}
	fmt.Println("Value of", tokens, "is :", EvalRPN(tokens))
}

func EvalRPN(tokens []string) int {
	op := "+-*/"

	s := Stack{list.New()}

	for _, v := range tokens {
		if !strings.Contains(op, v) {
			i, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
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
	return s.Pop().(int)
}
