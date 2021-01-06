/*
 * @lc app=leetcode.cn id=235 lang=golang
 *
 * [235] 二叉搜索树的最近公共祖先
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	lp := getPath(root, p)
// 	lq := getPath(root, q)
// 	var ancestor *TreeNode
// 	for i := 0; i < len(lp) && i < len(lq) && lp[i] == lq[i]; i++ {
// 		ancestor = lp[i]
// 	}
// 	return ancestor
// }
// @lc code=end

// func getPath(root, node *TreeNode) (res []*TreeNode) {
// 	for root != node {
// 		if node.Val > root.Val {
// 			res = append(res, root)
// 			root = root.Right
// 		} else {
// 			res = append(res, root)
// 			root = root.Left
// 		}
// 	}
// 	res = append(res, root)
// 	return res
// }

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}