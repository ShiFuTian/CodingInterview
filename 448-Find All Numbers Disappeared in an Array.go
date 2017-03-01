package leetcode

func findDisappearedNumbers(nums []int) (ans []int) {
	for _, k := range nums {
		if k < 0 { // abs(k)
			k = -k
		}
		if nums[k-1] > 0 { // mark
			nums[k-1] = -nums[k-1]
		}
	}
	for i, num := range nums {
		if num > 0 {
			ans = append(ans, i+1)
		}
	}
	return
}
