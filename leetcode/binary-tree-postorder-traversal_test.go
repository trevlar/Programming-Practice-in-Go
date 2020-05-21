package leetcode

import (
	"reflect"
	"testing"
)

func TestPostOrderTraversal(t *testing.T) {
	tree := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}

	actual := postorderTraversal(tree)
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Values abstracted from tree is incorrect, got: %v, want: %v.", actual, expected)
	}
}
