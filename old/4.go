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

type Word struct {
	val string
	no  int // Layer of tree(from 1)
	pre *Word
}

type Queue struct {
	slice []interface{}
}

func (q *Queue) Enqueue(v interface{}) {
	q.slice = append(q.slice, v)
}

func (q *Queue) Dequeue() interface{} {
	v := q.slice[0]
	q.slice = q.slice[1:]
	return v
}

func (q *Queue) Empty() bool {
	return len(q.slice) == 0
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
			log.Fatal("Invalid input.")
		}
	}

	queue := new(Queue)
	queue.Enqueue(&Word{start, 1, nil})

	for !queue.Empty() { // BFS
		word, ok := queue.Dequeue().(*Word)
		if !ok {
			log.Fatal("Wrong type.")
		}

		if word.val == end {
			re := make([]string, 0, word.no)
			for {
				re = append(re, word.val)
				if word.pre == nil {
					break
				}
				word = word.pre
			}
			res = append(res, re)
		}

		for i, v := range dict {
			if diffCount(word.val, v) == 1 {
				queue.Enqueue(&Word{v, word.no + 1, word})
				if v != end { //
					dict = append(dict[:i], dict[i+1:]...)
				}
			}
		}
	}
	return
}

func main() {
	fmt.Println(WordLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
}
