/*
 * @lc app=leetcode.cn id=297 lang=golang
 *
 * [297] 二叉树的序列化与反序列化
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

 type Codec struct {
	//l []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	// DFS 前序遍历递归
	//if root == nil {
	//	return "x"
	//}
	//return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)

	// BFS 迭代
	if root == nil {
		return ""
	}
	q := []*TreeNode{root}
	res := []string{}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node != nil {
			res = append(res, strconv.Itoa(node.Val))
			q = append(q, node.Left)
			q = append(q, node.Right)
		} else {
			res = append(res, "x")
		}
	}
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	//l := strings.Split(data, ",")
	//var ldeserialize func(arr *[]string) *TreeNode
	//ldeserialize = func(arr *[]string) *TreeNode {
	//	v := (*arr)[0]
	//	*arr = (*arr)[1:]
	//	fmt.Println("v: ", v, " arr:", *arr)
	//	if v == "x" {
	//		return nil
	//	}
	//	val, _ := strconv.Atoi(v)
	//	root := &TreeNode{
	//		Val: val,
	//	}
	//	root.Left = ldeserialize(arr)
	//	root.Right = ldeserialize(arr)
	//	ldeserialize(arr)
	//	return root
	//}
	//root := ldeserialize(&l)
	//return root

	if data == "" {
		return nil
	}
	l := strings.Split(data, ",")
	val, _ := strconv.Atoi(l[0])
	root := &TreeNode{Val: val}
	
	q := []*TreeNode{root}
	for cursor := 1; cursor < len(l);cursor += 2 {
		node := q[0]
		fmt.Println("val: ", node.Val, " cursor: ", cursor, " q:", q, " l:", l)
		q = q[1:]
		lval := l[cursor]
		if lval != "x" {
			v, _ := strconv.Atoi(lval)
			node.Left = &TreeNode{Val:v}
			q = append(q, node.Left)
		}
		rval := l[cursor+1]
		if rval != "x" {
			v, _ := strconv.Atoi(rval)
			node.Right = &TreeNode{Val:v}
			q = append(q, node.Right)
		}
		
	}
	return root
}
