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
	queue := make(chan *TreeNode, 0)

	if root == nil {
		return result
	}

	queue <- root
	for len(queue) > 0 {
		levelNum := len(queue)
		var subList []int
		for i := 0; i < levelNum; i++ {
			node = <-queue
			if node.Left != nil {
				queue <- node.Left
			}
			if node.Right != nil {
				queue <- node.Right
			}
			subList = append(subList, node.Val)
		}
		result = append(result, subList)
	}

	return result
}

