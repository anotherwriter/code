/**
https://leetcode.com/problems/intersection-of-two-linked-lists/

Given the heads of two singly linked-lists headA and headB, return the node at which the two lists intersect.
If the two linked lists have no intersection at all, return null.

Follow up: Could you write a solution that runs in O(m + n) time and use only O(1) memory?
*/
package main

import (
	"fmt"
	"math"

	"github.com/anotherwriter/code/leetcode/list"
)

func main() {
	r := list.BuildList([]int{8, 4, 5})

	r1 := list.BuildList([]int{4, 1})
	r2 := list.BuildList([]int{5, 6, 1})
	r1.Next.Next = r
	r2.Next.Next.Next = r

	fmt.Println(getIntersectionNode2(r1, r2))
}

// O(m+n) O(m)
func getIntersectionNode(headA, headB *list.ListNode) *list.ListNode {
	nodeMap := map[*list.ListNode]struct{}{}
	for n := headA; n != nil; n = n.Next {
		nodeMap[n] = struct{}{}
	}

	for n := headB; n != nil; n = n.Next {
		_, ok := nodeMap[n]
		if ok {
			return n
		}
	}
	return nil
}

// O(2m+2n) O(1)
func getIntersectionNode2(headA, headB *list.ListNode) *list.ListNode {
	var len1, len2 int
	for n := headA; n != nil; n = n.Next {
		len1++
	}
	for n := headB; n != nil; n = n.Next {
		len2++
	}

	n1, n2 := headA, headB
	tmpNode := &n1
	if len1 < len2 {
		tmpNode = &n2
	}
	for i := 1; i <= int(math.Abs(float64(len1-len2))); i++ {
		if tmpNode != nil {
			*tmpNode = (*tmpNode).Next
		}
	}

	for n1 != n2 {
		if n1 != nil {
			n1 = n1.Next
		}
		if n2 != nil {
			n2 = n2.Next
		}
	}

	return n1
}

// A walk path is listA + listB
// B walk path is listB + listA
// if has the intersection node, A will meet B
// refer: https://leetcode.cn/problems/intersection-of-two-linked-lists/solutions/811625/xiang-jiao-lian-biao-by-leetcode-solutio-a8jn/
func getIntersectionNode3(headA, headB *list.ListNode) *list.ListNode {
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}

		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
