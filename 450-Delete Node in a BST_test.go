package leetcode

import (
	"bytes"
	"strconv"
	"testing"
	"unsafe"
)

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewTree build a tree from an array of int.
// such as [5,3,6,2,4,-1,7]
// 			5
// 		3		6
// 	   2 4   nil 7
//func NewTree(input []int) (root *TreeNode) {
//	nodeList := make([]*TreeNode, len(input))
//	for i, val := range input {
//		if val != -1 {
//			nodeList[i] = &TreeNode{Val: val}
//		}
//	}
//	for i := 0; i < len(nodeList)/2; i++ {
//		node := nodeList[i]
//		if node != nil {
//			node.Left = nodeList[2*i+1]
//			node.Right = nodeList[2*i+2]
//		}
//	}
//	return nodeList[0]
//}
func NewTree(input []int) (root *TreeNode) {
	qin := &PointerQueue{}
	for _, val := range input {
		if val != -1 {
			qin.Enque(unsafe.Pointer(&TreeNode{Val: val}))
		} else {
			qin.Enque(nil)
		}
	}

	root = (*TreeNode)(qin.Deque())
	if root == nil {
		return nil
	}

	qout := &PointerQueue{}
	qout.Enque(unsafe.Pointer(root))

	for !qin.Empty() {
		node := (*TreeNode)(qout.Deque())
		if node != nil {
			node.Left = (*TreeNode)(qin.Deque())
			qout.Enque(unsafe.Pointer(node.Left))
			node.Right = (*TreeNode)(qin.Deque()) // qin maybe empty
			qout.Enque(unsafe.Pointer(node.Right))
		}
	}

	return root
}

func (root *TreeNode) String() string {
	var ret bytes.Buffer
	ret.WriteString("[")

	q := &PointerQueue{}
	q.Enque(unsafe.Pointer(root))

	for !q.Empty() {
		node := (*TreeNode)(q.Deque())
		ret.WriteString(strconv.Itoa(node.Val) + ",")

		if node.Left != nil {
			q.Enque(unsafe.Pointer(node.Left))
		}
		if node.Right != nil {
			q.Enque(unsafe.Pointer(node.Right))
		}
	}

	ret.WriteString("]")
	return ret.String()
}

func Test_deleteNode(t *testing.T) {
	type args struct {
		root *TreeNode
		key  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"1", args{NewTree([]int{5, 3, 6, 2, 4, -1, 7}), 3},
			NewTree([]int{5, 4, 6, 2, -1, -1, 7})},
		{"2", args{NewTree([]int{10, 4, 15, 2, -1, 13, 18, -1, 3, -1, 11, 14, 16, -1}), 10},
			NewTree([]int{13, 4, 15, 2, -1, 11, 18, -1, 3, -1, -1, 14, 16})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNode(tt.args.root, tt.args.key); got.String() != tt.want.String() {
				t.Errorf("deleteNode() = %s, want %s", got, tt.want)
			}
		})
	}
}
