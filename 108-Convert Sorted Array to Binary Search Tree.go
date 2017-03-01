package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 { // base case
		return nil
	}

	middle := len(nums) / 2

	node := &TreeNode{Val: nums[middle]}
	node.Left = sortedArrayToBST(nums[:middle])
	node.Right = sortedArrayToBST(nums[middle+1:])

	return node
}
