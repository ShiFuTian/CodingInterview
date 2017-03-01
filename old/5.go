// 5. Median of Two Sorted Arrays
// There are two sorted arrays A and B of size m and n respectively.
// Find the median of the two sorted arrays.
// The overall run time complexity should be O(log (m+n)).

package main

import (
	"fmt"
	"log"
)

func FindKth(A, B []int, K int) int {
	if K > len(A)+len(B)-1 {
		log.Fatal("K out of range.")
	}

	k := K
	aStart, aEnd := 0, len(A)-1
	bStart, bEnd := 0, len(B)-1

	for {
		aLen, bLen := aEnd-aStart+1, bEnd-bStart+1
		if aLen == 0 {
			return B[bStart+k]
		}
		if bLen == 0 {
			return A[aStart+k]
		}
		if k == 0 {
			if A[aStart] < B[bStart] {
				return A[aStart]
			}
			return B[bStart]
		}

		ak := k * aLen / (aLen + bLen)
		bk := k - 1 - ak // ak + bk + 1 = k is significant

		aMid := aStart + ak
		bMid := bStart + bk
		// (aMid -aStart) + (bMid - bStart) = k + 1
		// 0, 1, 2, 3, ... , k(th)

		if A[aMid] < B[bMid] {
			k -= (ak + 1)
			aStart = aMid + 1 // A[0]~A[aMid] can't be kth
			bEnd = bMid       // B[bMid+1]~B[bEnd] can't be kth
		} else {
			k -= (bk + 1)
			aEnd = aMid
			bStart = bMid + 1
		}
	}
}

func main() {
	A := []int{1, 3, 5, 7, 9}
	B := []int{2, 4, 6, 8, 10}
	for i := 0; i < 11; i++ {
		fmt.Println(FindKth(A, B, i))
	}
}
