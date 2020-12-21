/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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

 // 迭代
//  func isValidBST(root *TreeNode) bool {
// 	q := []*TreeNode{}
// 	pre := math.MinInt64
// 	for len(q) > 0 || root != nil {
// 		for root != nil {
// 			q = append(q, root)
// 			root = root.Left
// 		}
// 		root = q[len(q)-1]
// 		q = q[:len(q)-1]
// 		if root.Val <= pre {
// 			return false
// 		}
// 		pre = root.Val
// 		root = root.Right
// 	}
// 	return true
// }


func isValidBST(root *TreeNode) bool {
	var inorderTraversal func(root *TreeNode) bool
	pre := math.MinInt64
	inorderTraversal = func(root *TreeNode) bool{
		if root == nil {
			return true
		}
		if !inorderTraversal(root.Left) {
			return false
		}
		if root.Val <= pre {
			return false
		}
		pre = root.Val	
		if !inorderTraversal(root.Right) {
			return false
		}
		return true
	}
	return inorderTraversal(root)
}
