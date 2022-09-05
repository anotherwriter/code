package main

import "fmt"

/**
https://leetcode.com/problems/sort-colors/
https://leetcode.cn/problems/sort-colors/

Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

You must solve this problem without using the library's sort function.

Example 1:
	Input: nums = [2,0,2,1,1,0]
	Output: [0,0,1,1,2,2]

Example 2:
	Input: nums = [2,0,1]
	Output: [0,1,2]

Constraints:
 * n == nums.length
 * 1 <= n <= 300
 * nums[i] is either 0, 1, or 2.

Follow up: Could you come up with a one-pass algorithm using only constant extra space?
*/
func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors5(nums)
	fmt.Println(nums)

	nums = []int{2, 0, 1}
	sortColors5(nums)
	fmt.Println(nums)
}

// red(0) white(1) blue(2)
// based on swap sort
func sortColors1(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				// swap
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

// based on bubble sort
// much better than sortColors1(using leetcode test data)
func sortColors2(nums []int) {
	//var swapped bool
	for i := 0; i < len(nums)-1; i++ {
		//swapped = false
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				//swapped = true
				nums[j], nums[j+1] = nums[j+1], nums[j] // swap
			}
		}
		//if !swapped {
		//	return
		//}
	}
}

// based on counting num for each color
func sortColors3(nums []int) {
	var p0, p1, p2 int
	for _, v := range nums {
		switch v {
		case 0:
			p0++
		case 1:
			p1++
		case 2:
			p2++
		}
	}
	setNums(nums, 0, p0-1, 0)
	setNums(nums, p0, p0+p1-1, 1)
	setNums(nums, p0+p1, len(nums)-1, 2)
}

func setNums(nums []int, s, e, v int) {
	for ; s <= e; s++ {
		nums[s] = v
	}
}

// swap 0 to head of array and then swap 1 to the head of array
// Time Complexity: O(n)
// Extra Space: O(1)
func sortColors4(nums []int) {
	var p0 int
	for i, v := range nums {
		if v == 0 {
			nums[p0], nums[i] = 0, nums[p0]
			p0++
		}
	}

	var p1 int
	for i := p0; i < len(nums); i++ {
		if nums[i] == 1 {
			nums[p0+p1], nums[i] = 1, nums[p0+p1]
			p1++
		}
	}
}

// one-pass algorithm using only constant extra space
// double pointers
func sortColors5(nums []int) {
	var p0, p1 int

	for i := range nums {
		if nums[i] == 0 {
			nums[p0], nums[i] = 0, nums[p0]
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p0++
			p1++
		} else if nums[i] == 1 {
			nums[p1], nums[i] = 1, nums[p1]
			p1++
		}
	}
}
