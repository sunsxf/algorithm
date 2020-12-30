/*
 * @lc app=leetcode.cn id=701 lang=golang
 *
 * [701] 二叉搜索树中的插入操作
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
 func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	pre := &TreeNode{}
	child := "left"
	p := root
	for p != nil {

		pre = p
		if val > p.Val {
			p = p.Right
			child = "right"
		} else {
			p = p.Left
			child = "left"
		}
	}
	if child == "left" {
		pre.Left = &TreeNode{Val: val}
	} else {
		pre.Right = &TreeNode{Val: val}
	}
	return root
}

// @lc code=end

func insertIntoBST1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	p := root
	fmt.Println(p, root)
	for p != nil {
		if val > p.Val {
			if p.Right == nil {
				p.Right = &TreeNode{Val: val}
				break
			}
			p = p.Right

		} else {
			if p.Left == nil {
				p.Left = &TreeNode{Val: val}
				break
			}
			p = p.Left
		}
	}
	return root
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	
	} else {
		root.Left = insertIntoBST(root.Left, val)
	}
	
	return root
}
