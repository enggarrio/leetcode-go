package convertbsttogreatertree

// https://leetcode.com/problems/convert-bst-to-greater-tree/description/

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

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	getResult(root, &sum)
	return root
}
