/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(hight(root.Left)-hight(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func hight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(hight(root.Left), hight(root.Right)) + 1
}

func max(l, r int) int {
	if l > r {
		return l
	}
	return r
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
// @lc code=end

