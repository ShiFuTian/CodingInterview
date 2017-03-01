package leetcode

import (
	"strconv"
	"testing"
)

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(input []int) (output *ListNode) {
	output = &ListNode{Next: nil} // head, numpy node
	p := output
	for _, val := range input {
		p.Next = &ListNode{val, nil}
		p = p.Next
	}
	return output.Next // drop head
}

func (l *ListNode) Equal(ll *ListNode) bool {
	// p1, p2 := l, ll
	// for p1 != nil && p2 != nil {
	// 	if p1.Val != p2.Val {
	// 		return false
	// 	}
	// 	p1, p2 = p1.Next, p2.Next
	// }
	// if p1 != nil || p2 != nil {
	// 	return false
	// }
	// return true
	return l.String() == ll.String()
}

func (l *ListNode) String() string {
	var ret string
	for p := l; p != nil; p = p.Next {
		ret = ret + strconv.Itoa(p.Val) + ","
	}
	return ret
}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1", args{NewList([]int{1, 9, 9, 9}),
			NewList([]int{2, 3, 3, 3, 3, 2})},
			NewList([]int{2, 3, 5, 3, 3, 1}),
		},
		{"1", args{NewList([]int{0}),
			NewList([]int{0})},
			NewList([]int{0}),
		},
		{"1", args{NewList([]int{9, 9, 9, 9}),
			NewList([]int{1})},
			NewList([]int{1, 0, 0, 0, 0}),
		},
		{"1", args{NewList([]int{1, 9}),
			NewList([]int{2, 5, 9, 9})},
			NewList([]int{2, 6, 1, 8}),
		},
		{"1", args{NewList([]int{2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 9}),
			NewList([]int{5, 6, 4, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 9, 9, 9, 9})},
			NewList([]int{8, 0, 7, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 7, 2, 4, 3, 8}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !got.Equal(tt.want) {
				t.Errorf("addTwoNumbers() = %s want %s", got, tt.want)
			}
		})
	}
}
