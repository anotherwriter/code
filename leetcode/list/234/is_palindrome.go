package main

import (
	"fmt"
)

/**
https://leetcode.com/problems/palindrome-linked-list/
https://leetcode.cn/problems/palindrome-linked-list/

Given the head of a singly linked list, return true if it is a palindrome or false otherwise.

Example 1:
	Input: head = [1,2,2,1]
	Output: true

Example 2:
	Input: head = [1,2]
	Output: false

Constraints:
 * The number of nodes in the list is in the range [1, 10^5].
 * 0 <= Node.val <= 9

Follow up: Could you do it in O(n) time and O(1) space?
*/
func main() {
	root := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  -4,
					Next: nil,
				},
			},
		},
	}
	for n := root; n != nil; n = n.Next {
		fmt.Printf("%d ", n.Val)
	}
	fmt.Println("")

	fmt.Println(isPalindrome(root))
	fmt.Println(isPalindrome1(root))

	root.Next = nil
	fmt.Println(isPalindrome(root))
	fmt.Println(isPalindrome1(root))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// O(N) O(N)
func isPalindrome(head *ListNode) bool {
	var elems []int
	for n := head; n != nil; n = n.Next {
		elems = append(elems, n.Val)
	}

	for i, j := 0, len(elems)-1; i <= j; {
		if elems[i] != elems[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// O(N) O(1)
func isPalindrome1(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head != nil && head.Next == nil {
		return true
	}

	// get length of list
	var length int
	for tail := head; tail != nil; tail = tail.Next {
		length++
	}

	// split the list
	l2, r1 := length/2-1, length/2+1
	if length%2 == 0 {
		r1 = length / 2
		l2 = r1 - 1
	}

	var l2Node, r1Node *ListNode
	var i int
	for n := head; n != nil; n = n.Next {
		switch i {
		case l2:
			l2Node = n
		case r1:
			r1Node = n
			break
		}
		i++
	}

	// reverse list
	l2Node.Next = nil
	rHead := reverseList(r1Node)

	// compare
	n1, n2 := head, rHead
	for i := 0; i < length/2; i++ {
		if n1.Val != n2.Val {
			return false
		}
		n1 = n1.Next
		n2 = n2.Next
	}

	return true
}

// LeetCode
// https://leetcode.cn/problems/palindrome-linked-list/solution/hui-wen-lian-biao-by-leetcode-solution/
func isPalindromeLeetCode(head *ListNode) bool {
	if head == nil {
		return true
	}

	// 找到前半部分链表的尾节点并反转后半部分链表
	firstHalfEnd := endOfFirstHalf(head)
	secondHalfStart := reverseList(firstHalfEnd.Next)

	// 判断是否回文
	p1 := head
	p2 := secondHalfStart
	result := true
	for result && p2 != nil {
		if p1.Val != p2.Val {
			result = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	// 还原链表并返回结果
	firstHalfEnd.Next = reverseList(secondHalfStart)
	return result
}

func endOfFirstHalf(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode

	for n := head; n != nil; {
		next := n.Next

		n.Next = pre
		pre = n

		n = next
	}

	return pre
}
