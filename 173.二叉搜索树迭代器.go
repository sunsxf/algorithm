/*
 * @lc app=leetcode.cn id=173 lang=golang
 *
 * [173] 二叉搜索树迭代器
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
 
 type BSTIterator struct {
	bst *TreeNode
	q   []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		bst: root,
		q:   getTreeQueue(root),
	}
}

func (this *BSTIterator) Next() int {
	node := this.q[0]
	this.q = this.q[1:]
	return node.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.q) > 0
}

func getTreeQueue(root *TreeNode) (q []*TreeNode) {
	var inorderTraversal func(root *TreeNode)
	inorderTraversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorderTraversal(root.Left)
		q = append(q, root)
		inorderTraversal(root.Right)
	}
	inorderTraversal(root)
	return q
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end

