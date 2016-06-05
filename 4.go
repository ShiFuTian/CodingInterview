// 4. Word Ladder
// Given two words (start and end), and a dictionary,
// find the length of shortest transformation sequence
// from start to end, such that only one letter can be
// changed at a time and each intermediate word must
// exist in the dictionary. For example, given:
//
// start = "hit"
// end = "cog"
// dict = ["hot","dot","dog","lot","log"]
//
// One shortest transformation is
// "hit" -> "hot" -> "dot" -> "dog" -> "cog",
// the program should return its length 5.

package main

import (
	//"errors"
	"fmt"
	"log"
)

type Queue interface {
	Enqueue(v interface{})
	Dequeue() interface{}
	Empty() bool
}

type Word struct {
	val string
	no  int
	pre *Word
}

type WordQueue struct {
	slice []Word // Use pointer?
}

func (w *WordQueue) Enqueue(v interface{}) {
	w.slice = append(w.slice, v.(Word))
}

func (w *WordQueue) Dequeue() interface{} { // Use Base?
	v := w.slice[0]
	w.slice = w.slice[1:]
	return v
}

func (w *WordQueue) Empty() bool { // Use Base?
	return len(w.slice) == 0
}

func diffCount(s, t string) int {
	cnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			cnt++
		}
	}
	return cnt
}

func WordLadders(start, end string, dict []string) (res [][]string) {
	dict = append(dict, end)
	l := len(start)
	for _, v := range dict { // Check input
		if len(v) != l {
			log.Fatal("Invalid input!")
		}
	}

	queue := new(WordQueue)
	queue.Enqueue(Word{start, 1, nil})

	for !queue.Empty() { // BFS
		word := queue.Dequeue().(Word)
		if word.val == end {
			re := make([]string, 0, word.no)
			for {
				re = append(re, word.val)
				if word.pre == nil {
					break
				}
				word = *word.pre
			}
			res = append(res, re)
		}

		for i, v := range dict {
			if diffCount(word.val, v) == 1 {
				queue.Enqueue(Word{v, word.no + 1, &word})
				if v != end { //
					dict = append(dict[:i], dict[i+1:]...)
				}
			}
		}
	}
	return
}

func main() {
	fmt.Println(WordLadders("hit", "cog", []string{"hot", "dot", "dcg", "lot", "lcg"}))
}
