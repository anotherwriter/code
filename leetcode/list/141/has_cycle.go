package main

import "fmt"

/**
https://leetcode.com/problems/linked-list-cycle/
https://leetcode.cn/problems/linked-list-cycle/

Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by
continuously following the next pointer. Internally, pos is used to denote the index of the node that
tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.

Example 1:
	Input: head = [3,2,0,-4], pos = 1
	Output: true
	Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).

Example 2:
	Input: head = [1,2], pos = 0
	Output: true
	Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.

Example 3:
	Input: head = [1], pos = -1
	Output: false
	Explanation: There is no cycle in the linked list.

Constraints:
 * The number of the nodes in the list is in the range [0, 10^4].
 * -10^5 <= Node.val <= 10^5
 * pos is -1 or a valid index in the linked-list.

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
	fmt.Println(hasCycle2(root))
}

// O(N) O(N)
// hash map
func hasCycle(head *ListNode) bool {
	nodeMap := map[*ListNode]struct{}{}
	for n := head; n != nil; n = n.Next {
		if _, ok := nodeMap[n]; ok {
			return true
		}
		nodeMap[n] = struct{}{}
	}

	return false
}

// O(N) O(1)
// Floyd Cycle Detection Algorithm(Tortoise and Hare Algorithm)
// use double pointer: fast pointer, slow pointer
// - list doesn't have cycle:
//      fast ptr will arrive the tail before slow ptr. Each node can be accessed at most twice
// - list have cycle:
//      After each round, the distance between the fast and slow pointers will decrease by one.
//      The initial distance is the length of the ring, so at most N rounds are moved
func hasCycle2(head *ListNode) bool {
	slowPtr, fastPtr := head, head

	for fastPtr != nil && fastPtr.Next != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next

		if fastPtr == slowPtr {
			return true
		}
	}

	return false
}

func hasCycle2_2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next

	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
