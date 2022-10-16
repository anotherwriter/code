/**
https://leetcode.com/problems/remove-nth-node-from-end-of-list/

Given the head of a linked list, remove the nth node from the end of the list and return its head.

Example 1:
	Input: head = [1,2,3,4,5], n = 2
	Output: [1,2,3,5]

Example 2:
	Input: head = [1], n = 1
	Output: []

Example 3:
	Input: head = [1,2], n = 1
	Output: [1]

Constraints:
 * The number of nodes in the list is sz.
 * 1 <= sz <= 30
 * 0 <= Node.val <= 100
 * 1 <= n <= sz

Follow up: Could you do this in one pass?
*/

package main

import "github.com/anotherwriter/code/leetcode/list"

func main() {
	cases := [][]int{
		{1, 2, 3, 4, 5},
	}

	for _, nums := range cases {
		root := list.BuildList(nums)
		list.PrintList(root)

		root = removeNthFromEnd(root, 2)
		list.PrintList(root)
	}
}

// conquer based on compute len of list
// O(N) O(1)
func removeNthFromEnd(head *list.ListNode, n int) *list.ListNode {
	// get len of list
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}

	targetIndex := length - n
	var i int
	var preNode *list.ListNode
	for node := head; node != nil; node = node.Next {
		if i == targetIndex {
			if preNode == nil {
				return node.Next
			}
			preNode.Next = node.Next
		}

		i++
		preNode = node
	}

	return head
}

// other solutions
// refers: https://leetcode.cn/problems/remove-nth-node-from-end-of-list/solutions/450350/shan-chu-lian-biao-de-dao-shu-di-nge-jie-dian-b-61/
// 1.double pointers
// 	first ptr == head; second ptr == head walk n nodes
// 	when second ptr is nil, first prt is the deleted node
