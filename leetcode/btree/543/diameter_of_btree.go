/**
https://leetcode.com/problems/diameter-of-binary-tree/

Given the root of a binary tree, return the length of the diameter of the tree.

The diameter of a binary tree is the length of the longest path between any two nodes in a tree.
This path may or may not pass through the root.

The length of a path between two nodes is represented by the number of edges between them.

Example 1:
	Input: root = [1,2,3,4,5]
	Output: 3
	Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].
*/
package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/diameter-of-binary-tree/solution/er-cha-shu-de-zhi-jing-by-leetcode-solution/
func diameterOfBinaryTree(root *TreeNode) int {
	ans := 1
	depth(&ans, root)

	return ans - 1
}

func depth(ans *int, node *TreeNode) int {
	if node == nil {
		return 0 // 访问到空节点了，返回0
	}
	L := depth(ans, node.Left)  // 左儿子为根的子树的深度
	R := depth(ans, node.Right) // 右儿子为根的子树的深度

	*ans = max(*ans, L+R+1) // 计算d_node即L+R+1 并更新ans
	return max(L, R) + 1    // 返回该节点为根的子树的深度
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
