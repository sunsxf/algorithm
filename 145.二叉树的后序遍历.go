/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
 */

// @lc code=start
// /**
//  * Definition for a binary tree node.
//  * type TreeNode struct {
//  *     Val int
//  *     Left *TreeNode
//  *     Right *TreeNode
//  * }
//  */

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil 
	}
	stack1 := []*TreeNode{root}
	stack2 := []*TreeNode{}
	res := []int{}
	for len(stack1) > 0 {
		node := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]
		stack2 = append(stack2, node)
		if node.Left != nil {
			stack1 = append(stack1,node.Left)
		}
		if node.Right != nil {
			stack1 = append(stack1,node.Right)
		}
	}
	for len(stack2) > 0 {
		node := stack2[len(stack2)-1]
		stack2 = stack2[:len(stack2)-1]
		res = append(res, node.Val)
	}
	return res
}
// @lc code=end

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil 
	}
	stack, cur, prev, res := []*TreeNode{}, root, *TreeNode{}, []int{}
	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		if len(stack) > 0 {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if cur.Right == nil || cur.Right == prev {
				// 右节点为空或者右节点等于prev，打印当前节点，preview=当前节点，cur=nil,不继续压栈
				res = append(res, cur.Val)
				prev = cur
				cur = nil
			} else {
				// 右节点不为空且右节点不和previwe相同，表示右子树还未遍历，继续后续遍历右子树，先不打印，继续压栈
				cur = cur.Right
				stack = append(stack, cur)
			}
		}
	}
	return res
}
// @lc code=end