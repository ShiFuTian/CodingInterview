package leetcode

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tt := struct {
		name string
		args args
		want float64
	}{
		"1", args{
			nums1: []int{1, 3, 5, 7, 9},
			nums2: []int{2, 4, 6, 8, 10},
		}, 5.5,
	}
	for i := 1; i <= 10; i++ {
		if got := findKth(tt.args.nums1, tt.args.nums2, i); got != i {
			t.Errorf("findKth(%v) = %v, want %v", i, got, i)
		}
	}
	if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
		t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
	}

}
