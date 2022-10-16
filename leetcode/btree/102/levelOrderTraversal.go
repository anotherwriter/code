package main

import (
	"container/list"
	"fmt"
)

/**
https://leetcode.com/problems/binary-tree-level-order-traversal/
Desc:
Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

Case1
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	fmt.Println(levelOrder(root))
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int

	l := list.List{}
	l.PushBack(root)

	for l.Len() != 0 {
		var values []int

		levelSize := l.Len()
		for i := 0; i < levelSize; i++ {
			e := l.Front()
			l.Remove(e)

			node := e.Value.(*TreeNode)
			values = append(values, node.Val)

			if node.Left != nil {
				l.PushBack(node.Left)
			}
			if node.Right != nil {
				l.PushBack(node.Right)
			}
		}

		result = append(result, values)
	}

	return result
}
