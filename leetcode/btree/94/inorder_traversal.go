/**
https://leetcode.com/problems/binary-tree-inorder-traversal/

Given the root of a binary tree, return the inorder traversal of its nodes' values.

Example 1:
	Input: root = [1,null,2,3]
	Output: [1,3,2]

Example 2:
	Input: root = []
	Output: []

Example 3:
	Input: root = [1]
	Output: [1]

Constraints:
 * The number of nodes in the tree is in the range [0, 100].
 * -100 <= Node.val <= 100
*/
package main

import (
	"fmt"
)

func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}

	fmt.Println(inorderTraversal(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var nums []int
	var stack []*TreeNode

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		nums = append(nums, root.Val)
		root = root.Right
	}

	return nums
}

func inorderTraversal1(root *TreeNode) []int {
	var nums []int
	recurse(root, &nums)
	return nums
}

func recurse(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	recurse(root.Left, nums)
	*nums = append(*nums, root.Val)
	recurse(root.Right, nums)
}
