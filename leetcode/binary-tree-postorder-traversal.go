/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) (result []int) {
	if root == nil {
		return result
	}

	var stack []*TreeNode
	stack = append(stack, root)

	return recursivePreorderTraversal(stack, result)
}

func recursivePreorderTraversal(stack []*TreeNode, result []int) []int {
	node, stack := stack[len(stack)-1], stack[:len(stack)-1]
	result = append([]int{node.Val}, result...)

	if node.Left != nil {
		stack = append(stack, node.Left)
	}

	if node.Right != nil {
		stack = append(stack, node.Right)
	}

	if len(stack) > 0 {
		return recursivePreorderTraversal(stack, result)
	} else {
		return result
	}
}