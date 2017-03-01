package leetcode

func singleNumber2(nums []int) int {
	ones, twos := 0, 0
	for _, num := range nums {
		ones ^= (num & ^twos)
		twos ^= (num & ^ones)
	}
	return ones
}
