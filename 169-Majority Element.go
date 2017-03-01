package leetcode

import (
	"math/rand"
	"sort"
	"time"
)

// "Boyer-Moore Majority Vote Algorithm"
// Genius! O(n) time, O(1) space
func majorityElement(nums []int) int {
	major := nums[0] // assume that the array is non-empty
	count := 1
	for _, num := range nums[1:] {
		if major == num {
			count++ // IMP: major is reinforced++
		} else {
			count-- // IMP: major is counteracted--
		}
		if count == 0 { // IMP: need a new major
			major = num
			count = 1
		}
	}
	return major // and the majority element always exist in the array.
}

// ATTEMPT 1 DIVIDE & CONQUER
// cut nums into 3 sub nums, however I cannot
// solve the base case of len(nums) == 2.

// ATTEMPT 2 MAP
// O(n) time, O(n) space
func majorityElement2(nums []int) int {
	count := make(map[int]int)
	for _, v := range nums {
		count[v]++
		if count[v] > len(nums)/2 {
			// You may assume that the array is non-empty,
			// and the majority element always exist in the array.
			return v
		}
	}
	// There is no majority element,
	// return nums[0] as a false result.
	return nums[0]
}

// PRIVATE 3 SORT
// O(n*logn) time, O(1) space
func majorityElement3(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// PRIVATE 4 RANDOMIZATION
func majorityElement4(nums []int) int {
	rand.Seed(time.Now().UnixNano())
	nlen := len(nums)
	for {
		rdm := nums[rand.Intn(nlen)]
		count := 0
		for _, num := range nums {
			if rdm == num {
				count++
			}
			if count > nlen/2 {
				return rdm
			}
		} // range
	}
}

// PRIVATE 5 DIVIDE & CONQUER
func majorityElement5(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	// divide
	mid := len(nums) / 2
	lm := majorityElement5(nums[:mid])
	rm := majorityElement5(nums[mid:])

	// conquer
	if lm == rm {
		return lm
	}
	// lm != rm
	var lmCnt, rmCnt int
	for _, v := range nums {
		if v == lm {
			lmCnt++
		} else if v == rm {
			rmCnt++
		}
	}
	if lmCnt > rmCnt {
		return lm
	}
	return rm
}

// PRIVATE 6 BIT MANIPULATION
// Another nice idea!
// The key lies in how to count the number of 1's on a specific bit.
// Specifically, you need a mask with a 1 on the i-th bit and 0 otherwise
// to get the i-th bit of each element in nums. The code is as follows.
func majorityElement6(nums []int) int {
	var major int
	nlen := len(nums)

	for i, mask := 0, 1; i < 32; i, mask = i+1, mask<<1 {
		bitCount := 0
		for j := 0; j < nlen; j++ {
			if nums[j]&mask > 0 {
				bitCount++
			}
			if bitCount > nlen/2 {
				major |= mask
				break
			}
		} // j
	} // i

	return major
}
