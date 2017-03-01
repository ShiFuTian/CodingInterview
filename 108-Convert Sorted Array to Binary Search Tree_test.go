package leetcode

import "testing"

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"1", args{[]int{1, 6, 9, 12, 100, 102}},
			NewTree([]int{12, 6, 102, 1, 9, 100})},
		{"2", args{[]int{1, 6, 9, 12, 100}},
			NewTree([]int{9, 6, 100, 1, -1, 12})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrayToBST(tt.args.nums); got.String() != tt.want.String() {
				t.Errorf("deleteNode() = %s, want %s", got, tt.want)
			}
		})
	}
}
