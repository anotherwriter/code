/**
https://leetcode.com/problems/maximum-depth-of-binary-tree/

Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path
from the root node down to the farthest leaf node.
*/
package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// based on bfs
func maxDepth(root *TreeNode) int {
	return 0
}

// based on dfs
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right)))) + 1
}
