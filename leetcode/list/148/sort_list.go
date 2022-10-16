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
	"fmt"

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

		head = sortList(head)
		list.PrintList(head)
		fmt.Println(getListMidNode(head, nil))
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
	_sortList2(head, nil)
	return head
}

func _sortList2(head, tail *list.ListNode) {
	if head != nil {
		mid := getListMidNode(head, tail)
		_sortList2(head, mid)
		_sortList2(mid.Next, tail)
		mergeList2(head, mid, tail)
	}
}

func mergeList2(head, mid, tail *list.ListNode) {

}

func merge(head1, head2 *list.ListNode) *list.ListNode {
	dummyHead := &list.ListNode{}
	temp, n1, n2 := dummyHead, head1, head2

	for n1 != nil && n2 != nil {
		if n1.Val <= n2.Val {
			temp.Next = n1
			n1 = n1.Next
		} else {
			temp.Next = n2
			n2 = n2.Next
		}
		temp = temp.Next
	}
	if n1 != nil {
		temp.Next = n1
	} else if n2 != nil {
		temp.Next = n2
	}
	return dummyHead.Next
}

// note: in fact, it's mid.Next
func getListMidNode(head, tail *list.ListNode) *list.ListNode {
	if head == nil || head.Next == tail {
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}

	return slow
}
