package leetcode

func firstMissingPositive2(nums []int) int {
	k, n := 0, len(nums)
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			if k != i {
				nums[i], nums[k] = nums[k], nums[i]
			}
			k++
		}
	}

	for i := 0; i < k; i++ {
		pos := nums[i]
		if pos < 0 {
			pos = -pos
		}
		if pos <= k && nums[pos-1] > 0 {
			nums[pos-1] = -nums[pos-1]
		}
	}

	for i := 0; i < k; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}

	return k + 1
}
