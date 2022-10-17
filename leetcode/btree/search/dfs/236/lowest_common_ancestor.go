/**
https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/

Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.
According to the definition of LCA on Wikipedia:
	“The lowest common ancestor is defined between two nodes p and q as the lowest node in T
	that has both p and q as descendants (where we allow a node to be a descendant of itself).”

Example 1:
	Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
	Output: 3
	Explanation: The LCA of nodes 5 and 1 is 3.
*/
package main

import (
	"container/list"
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(v int) *TreeNode {
	if v == math.MinInt32 {
		return nil
	}
	return &TreeNode{Val: v}
}

func main() {
	root := buildBinaryTree([]int{3, 5, 1, 6, 2, 0, 8, math.MinInt, math.MinInt, 7, 4})
	fmt.Println(lowestCommonAncestor(root, root.Left, root.Right))
}

// dfs
// O(N) O(N)
// set x is the lowest common ancestor. x should satisfy one condition below:
// 1. x.left, x.right contain p, q
// 2. (x == p or x == q) and x.left or x.right contain p/q
// refer: https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/solutions/238552/er-cha-shu-de-zui-jin-gong-gong-zu-xian-by-leetc-2/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}

	return left
}

func buildBinaryTree(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	root := &TreeNode{Val: values[0]}

	j := 1
	l := list.List{}
	l.PushBack(root)
	for l.Len() != 0 {
		//levelSize := l.Len()
		//for i := 0; i < levelSize; i++ {
		e := l.Front()
		l.Remove(e)

		node := e.Value.(*TreeNode)
		if j < len(values) {
			node.Left = NewTreeNode(values[j])
			l.PushBack(node.Left)
			j++
		}
		if j < len(values) {
			node.Right = NewTreeNode(values[j])
			l.PushBack(node.Right)
			j++
		}
		//}
	}

	return root
}
