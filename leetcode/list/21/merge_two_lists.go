/**
https://leetcode.com/problems/merge-two-sorted-lists/

You are given the heads of two sorted linked lists list1 and list2.
Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.
Return the head of the merged linked list.

Example 1:
	Input: list1 = [1,2,4], list2 = [1,3,4]
	Output: [1,1,2,3,4,4]

Example 2:
	Input: list1 = [], list2 = []
	Output: []

Example 3:
	Input: list1 = [], list2 = [0]
	Output: [0]

Constraints:
 * The number of nodes in both lists is in the range [0, 50].
 * -100 <= Node.val <= 100
 * Both list1 and list2 are sorted in non-decreasing order.
*/
package main

import (
	"fmt"

	"github.com/anotherwriter/code/leetcode/list"
)

func main() {
	cases := [][][]int{
		{{1, 2, 4}, {1, 3, 4}},
		{{}, {}},
		{{}, {0}},
	}
	for _, numsArray := range cases {
		fmt.Println("------------------")
		l1, l2 := list.BuildList(numsArray[0]), list.BuildList(numsArray[1])
		list.PrintList(l1)
		list.PrintList(l2)
		r := mergeTwoLists(l1, l2)
		list.PrintList(r)
	}
}

func mergeTwoLists(list1 *list.ListNode, list2 *list.ListNode) *list.ListNode {
	n1, n2 := list1, list2

	tmpNode := new(list.ListNode)
	node := tmpNode
	for ; n1 != nil && n2 != nil; node = node.Next {
		if n1.Val < n2.Val {
			node.Next = n1
			n1 = n1.Next
		} else {
			node.Next = n2
			n2 = n2.Next
		}
	}

	if n1 != nil {
		node.Next = n1
	}
	if n2 != nil {
		node.Next = n2
	}

	return tmpNode.Next
}
