package main

import (
	"fmt"
)

/**
https://leetcode.com/problems/reverse-linked-list/
https://leetcode.cn/problems/reverse-linked-list/

Given the head of a singly linked list, reverse the list, and return the reversed list.

Example 1:
	Input: head = [1,2,3,4,5]
	Output: [5,4,3,2,1]

Example 2:
	Input: head = [1,2]
	Output: [2,1]

Example 3:
	Input: head = []
	Output: []

Constraints:
 * The number of nodes in the list is the range [0, 5000].
 * -5000 <= Node.val <= 5000

Follow up: A linked list can be reversed either iteratively or recursively. Could you implement both?
*/
func main() {
	root := buildList([]int{3, 2, 0, -4})
	printList(root)

	r := reverseList(root)
	printList(r)

	r2 := reverseList2(r)
	printList(r2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// iteratively
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for curr := head; curr != nil; {
		next := curr.Next

		curr.Next = prev

		prev = curr
		curr = next
	}

	return prev
}

// recursively
func reverseList2(head *ListNode) *ListNode {
	return reverseRecursively(nil, head)
}

func reverseRecursively(pre, curr *ListNode) *ListNode {
	if curr != nil {
		newPre := reverseRecursively(curr, curr.Next)
		curr.Next = pre

		pre = newPre
	}

	return pre
}

// https://leetcode.cn/problems/reverse-linked-list/solutions/551596/fan-zhuan-lian-biao-by-leetcode-solution-d1k2/
func reverseListFromLC(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseListFromLC(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	root := &ListNode{Val: nums[0]}

	node := root
	for i := 1; i < len(nums); i++ {
		node.Next = &ListNode{Val: nums[i]}
		node = node.Next
	}

	return root
}

func printList(root *ListNode) {
	var values []int
	for n := root; n != nil; n = n.Next {
		values = append(values, n.Val)
	}
	fmt.Println(values)
}
