package leetcode

import (
	"reflect"
	"testing"
)

func Test_findDisappearedNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{"1", args{[]int{4, 3, 2, 7, 8, 2, 3, 1}}, []int{5, 6}},
		{"1", args{[]int{4, 5, 2, 7, 8, 2, 3, 1}}, []int{6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findDisappearedNumbers(tt.args.nums); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("findDisappearedNumbers() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
