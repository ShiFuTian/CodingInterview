package leetcode

func singleNumber3(nums []int) []int {
	var diff int
	for _, num := range nums {
		diff ^= num
	}

	diff &= -diff

	var a, b int
	for _, num := range nums {
		if num&diff != 0 {
			a ^= num
		} else {
			b ^= num
		}
	}

	return []int{a, b}
}
