package main

import "fmt"

/**
https://leetcode.com/problems/linked-list-cycle-ii/
https://leetcode.cn/problems/linked-list-cycle-ii/

Given the head of a linked list, return the node where the cycle begins. If there is no cycle, return null.

  There is a cycle in a linked list if there is some node in the list that can be reached again
by continuously following the next pointer. Internally, pos is used to denote the index of the node that
tail's next pointer is connected to (0-indexed). It is -1 if there is no cycle. Note that pos is not passed as a parameter.

Do not modify the linked list.

Example 1:
	Input: head = [3,2,0,-4], pos = 1
	Output: tail connects to node index 1
	Explanation: There is a cycle in the linked list, where tail connects to the second node.

Example 2:
	Input: head = [1,2], pos = 0
	Output: tail connects to node index 0
	Explanation: There is a cycle in the linked list, where tail connects to the first node.

Example 3:
	Input: head = [1], pos = -1
	Output: no cycle
	Explanation: There is no cycle in the linked list.

Follow up: Can you solve it using O(1) (i.e. constant) memory?
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

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
	root.Next.Next.Next.Next = root.Next
	fmt.Println("cycle entry node:", detectCycle2(root))
}

func detectCycle(head *ListNode) *ListNode {
	nodeMap := map[*ListNode]struct{}{}
	for n := head; n != nil; n = n.Next {
		if _, ok := nodeMap[n]; ok {
			return n
		}
		nodeMap[n] = struct{}{}
	}

	return nil
}

func detectCycle2(head *ListNode) *ListNode {
	getMeetNode := func() *ListNode {
		slowPtr, fastPtr := head, head

		for fastPtr != nil && fastPtr.Next != nil {
			slowPtr = slowPtr.Next
			fastPtr = fastPtr.Next.Next

			if fastPtr == slowPtr {
				return fastPtr
			}
		}
		return nil
	}

	meetPtr := getMeetNode()
	if meetPtr == nil {
		return nil
	}

	ptr1, ptr2 := head, meetPtr
	for ptr1 != ptr2 {
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}
	return ptr1
}

func detectCycle3(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
