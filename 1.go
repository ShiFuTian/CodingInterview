// 1. Rotate Array
// You may have been using Java for a while. Do you think a simple Java array question
// can be a challenge? Let’s use the following problem to test.
// Problem: Rotate an array of n elements to the right by k steps.
// For example, with n = 7 and k = 3, the array [1,2,3,4,5,6,7] is rotated to [5,6,7,1,2,3,4].
// How many different ways do you know to solve this problem?
package main

import (
	"fmt"
	"log"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(nums)
	nums = rotateStraight(nums, 3)
	fmt.Println(nums)
	rotateBubble(nums, 3)
	fmt.Println(nums)
	rotateReversal(nums, 3)
	fmt.Println(nums)
}

// O(n) time, O(n) space
func rotateStraight(nums []int, k int) []int {
	if nums == nil || k < 0 {
		log.Fatal("Illegal argument!")
	}

	l := len(nums)
	nums2 := append(nums[l-k:], nums[:l-k]...)

	return nums2
}

// O(n*k) time, O(1) space
// This solution is like a bubble sort.
func rotateBubble(nums []int, k int) {
	if nums == nil || k < 0 {
		log.Fatal("Illegal argument!")
	}

	l := len(nums)
	for i := l - k; i < l; i++ {
		for j := i; j > i-l+k; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}

// O(n) time, O(1) space
// 1. Divide the array two parts: 1,2,3,4 and 5, 6
// 2. Rotate first part: 4,3,2,1,5,6
// 3. Rotate second part: 4,3,2,1,6,5
// 4. Rotate the whole array: 5,6,1,2,3,4
func rotateReversal(nums []int, k int) {
	if nums == nil || k < 0 {
		log.Fatal("Illegal argument!")
	}

	l := len(nums)
	m := l - k
	reverse(nums, 0, m-1)
	reverse(nums, m, l-1)
	reverse(nums, 0, l-1)
}

func reverse(nums []int, left int, right int) {
	if nums == nil || len(nums) == 1 {
		return
	}

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
