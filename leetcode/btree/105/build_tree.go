/**
https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree
and inorder is the inorder traversal of the same tree, construct and return the binary tree.

Example 1:
	Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
	Output: [3,9,20,null,null,15,7]
Example 2:
	Input: preorder = [-1], inorder = [-1]
	Output: [-1]

Constraints:
 * 1 <= preorder.length <= 3000
 * inorder.length == preorder.length
 * -3000 <= preorder[i], inorder[i] <= 3000
 * preorder and inorder consist of unique values.
 * Each value of inorder also appears in preorder.
 * preorder is guaranteed to be the preorder traversal of the tree.
 * inorder is guaranteed to be the inorder traversal of the tree.
*/
package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// O(n) O(n)
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	i := 0 // index of root in inorder
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	leftChildren := inorder[:i]
	rightChildren := inorder[i+1:]
	leftChildCnt := len(leftChildren)

	root.Left = buildTree(preorder[1:leftChildCnt+1], leftChildren)
	root.Right = buildTree(preorder[leftChildCnt+1:], rightChildren)
	return root
}
