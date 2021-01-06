/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}
// @lc code=end

func helper(nums []int, low, up int) *TreeNode {
	if low > up {
		return nil
	}
	mid := (low + up)/2
	root := &TreeNode{Val:nums[mid]}
	root.Left = helper(nums, low, mid-1)
	root.Right = helper(nums, mid+1, up)
	return root
}