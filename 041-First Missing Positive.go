package leetcode

func firstMissingPositive1(nums []int) int {
	n := len(nums)

	for _, k := range nums {
		for k >= 1 && k <= n { // place k in nums[k-1]
			if nums[k-1] != k {
				nums[k-1], k = k, nums[k-1]
			} else {
				break
			}
		}
	}

	for i, k := range nums { // search
		if k != i+1 {
			return i + 1
		}
	}

	return n + 1
}
