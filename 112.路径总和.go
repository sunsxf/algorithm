/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
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

// 广度优先搜索
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	nq, vq := []*TreeNode{root}, []int{root.Val}
	for len(nq) > 0 {
		node := nq[0]
		nq = nq[1:]
		val := vq[0]
		vq = vq[1:]
		if node.Left == nil && node.Right == nil && val == sum {
			return true
		}
		if node.Left != nil {
			nq = append(nq, node.Left)
			vq = append(vq, val + node.Left.Val)
		}
		if node.Right != nil {
			nq = append(nq, node.Right)
			vq = append(vq, val + node.Right.Val)
		}
	}
	return false
}
// @lc code=end

// 递归
// func recursiveHasPathSum(root *TreeNode, sum int) bool {
// 	if root == nil {
// 		return false
// 	}
// 	if root.Left == nil && root.Right == nil {
// 		return sum == root.Val
// 	}
// 	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
// }
//