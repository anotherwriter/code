/**
https://leetcode.com/problems/sort-list/
https://leetcode.cn/problems/sort-list/

Given the head of a linked list, return the list after sorting it in ascending order.

Example 1:
	Input: head = [4,2,1,3]
	Output: [1,2,3,4]

Example 2:
	Input: head = [-1,5,3,4,0]
	Output: [-1,0,3,4,5]

Example 3:
	Input: head = []
	Output: []

Constraints:
 * The number of nodes in the list is in the range [0, 5 * 10^4].
 * -10^5 <= Node.val <= 10^5

Follow up: Can you sort the linked list in O(n logn) time and O(1) memory (i.e. constant space)?
*/
package main

import (
	"github.com/anotherwriter/code/leetcode/list"
)

func main() {
	cases := [][]int{
		{4, 2, 1, 3},
		{-1, 5, 3, 4, 0},
		{},
	}
	for _, nums := range cases {
		head := list.BuildList(nums)
		list.PrintList(head)

		head = sortList3(head)
		list.PrintList(head)
	}
}

// based on swap sort
// O(N^2)
func sortList(head *list.ListNode) *list.ListNode {
	if head == nil {
		return nil
	}

	for n1 := head; n1.Next != nil; n1 = n1.Next {
		for n2 := n1.Next; n2 != nil; n2 = n2.Next {
			if n1.Val > n2.Val {
				n1.Val, n2.Val = n2.Val, n1.Val
			}
		}
	}

	return head
}

// based on top-to-bottom merge sort
// refers: https://leetcode.cn/problems/sort-list/solutions/492301/pai-xu-lian-biao-by-leetcode-solution/
func sortList2(head *list.ListNode) *list.ListNode {
	return _sortList2(head, nil)
}

func _sortList2(head, tail *list.ListNode) *list.ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil // specially
		return head
	}

	// get middle node
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	mid := slow
	root := mergeList2(_sortList2(head, mid), _sortList2(mid, tail))
	return root
}

func mergeList2(head1, head2 *list.ListNode) *list.ListNode {
	tmpHead := new(list.ListNode)
	node, n1, n2 := tmpHead, head1, head2

	for n1 != nil && n2 != nil {
		if n1.Val <= n2.Val {
			node.Next = n1
			n1 = n1.Next
		} else {
			node.Next = n2
			n2 = n2.Next
		}
		node = node.Next
	}

	if n1 != nil {
		node.Next = n1
	}
	if n2 != nil {
		node.Next = n2
	}

	return tmpHead.Next
}

// based on bottom-to-top merge sort
// refers: https://leetcode.cn/problems/sort-list/solutions/492301/pai-xu-lian-biao-by-leetcode-solution/
func sortList3(head *list.ListNode) *list.ListNode {
	if head == nil {
		return head
	}

	length := 0
	for n := head; n != nil; n = n.Next {
		length++
	}

	tmpHead := &list.ListNode{Next: head}
	for subLength := 1; subLength < length; subLength <<= 1 {
		prev, cur := tmpHead, tmpHead.Next
		for cur != nil {
			head1 := cur
			for i := 1; i < subLength && cur.Next != nil; i++ {
				cur = cur.Next
			}

			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}

			var next *list.ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}

			prev.Next = mergeList2(head1, head2)

			for prev.Next != nil {
				prev = prev.Next
			}
			cur = next
		}
	}
	return tmpHead.Next
}
