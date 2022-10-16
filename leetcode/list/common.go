package list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildList(nums []int) *ListNode {
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

func PrintList(root *ListNode) {
	var values []int
	for n := root; n != nil; n = n.Next {
		values = append(values, n.Val)
	}
	fmt.Println(values)
}
