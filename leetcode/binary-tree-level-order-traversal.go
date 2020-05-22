package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) (result [][]int) {
	var node *TreeNode

	if root == nil {
		return result
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		levelNum := len(queue)
		var subList []int
		for i := 0; i < levelNum; i++ {
			node, queue = queue[0], queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			subList = append(subList, node.Val)
		}
		result = append(result, subList)
	}

	return result
}
