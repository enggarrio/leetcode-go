package binarysearchtreetogreatersumtree

// https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/description/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// delete this before submit
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getResult(node *TreeNode, sum *int) {
	if node == nil {
		return
	}
	getResult(node.Right, sum)
	*sum += node.Val
	node.Val = *sum
	getResult(node.Left, sum)
}

func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	getResult(root, &sum)
	return root
}
