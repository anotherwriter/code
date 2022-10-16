package main

import (
	"container/list"
	"fmt"
	"math"
)

/**
https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/

Desc:
Given the root of a binary tree, the value of a target node target, and an integer k, return an array of the values of all nodes that have a distance k from the target node.

You can return the answer in any order.

Case1:
	Input: root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, k = 2
	Output: [7,4,1]
	Explanation: The nodes that are a distance 2 from the target node (with value 5) have values 7, 4, and 1.

Case2:
	Input: root = [1], target = 1, k = 3
	Output: []

Constraints:
	The number of nodes in the tree is in the range [1, 500].
	0 <= Node.val <= 500
	All the values Node.val are unique.
	target is the value of one of the nodes in the tree.
	0 <= k <= 1000
*/

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
	values := []int{3, 5, 1, 6, 2, 0, 8, math.MinInt32, math.MinInt32, 7, 4}

	root := buildBinaryTree(values)

	//fmt.Println(root)
	fmt.Println(distanceK(root, NewTreeNode(5), 2))
	//dfs(root)
	//bfs(root)
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	if k == 0 {
		return []int{target.Val}
	}

	layers, value2Layer := bfs(root)

	targetLayer := value2Layer[target.Val]

	var result []int
	if values, ok := layers[targetLayer+k]; ok {
		result = append(result, values...)
	}

	if targetLayer-k > 0 {
		result = append(result, layers[targetLayer-k]...)
	} else {

	}

	return result
}

// breadth first search
func bfs(root *TreeNode) (map[int][]int, map[int]int) {
	layers := map[int][]int{}    // layer => node values
	value2Layer := map[int]int{} // node value => layer

	queue := list.List{}
	queue.PushBack(root)

	layer := 1
	for queue.Len() != 0 {
		width := queue.Len()
		for j := 0; j < width; j++ {
			e := queue.Front()
			queue.Remove(e)
			node := e.Value.(*TreeNode)

			layers[layer] = append(layers[layer], node.Val)
			value2Layer[node.Val] = layer

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		layer++
	}

	//fmt.Println(layers)
	return layers, value2Layer
}

// depth first search
func dfs(root *TreeNode) {
	fmt.Println(root.Val)

	if root.Left != nil {
		dfs(root.Left)
	}

	if root.Right != nil {
		dfs(root.Right)
	}
}

func search(root *TreeNode, target *TreeNode, k int, values []int) []int {
	if k == 0 {
		return append(values, root.Val)
	}

	if target.Val == root.Val {
		search(target.Left, target, k-1, values)
		fmt.Println(root.Val)
		search(target.Right, target, k-1, values)
	}

	return values
}

func buildBinaryTree(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	root := NewTreeNode(values[0])
	l := list.List{}
	l.PushBack(root)

	j := 1
	for l.Len() != 0 {
		levelSize := l.Len()
		for i := 0; i < levelSize; i++ {
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
		}
	}

	return root
}
