package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) (result []int) {
	if root == nil {
		return
	}
	var stack []*TreeNode
	node := root

	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node, stack = stack[len(stack)-1], stack[:len(stack)-1]
		result = append(result, node.Val)
		node = node.Right
	}

	return
}
