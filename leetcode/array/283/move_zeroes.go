package main

import "fmt"

/**
https://leetcode.com/problems/move-zeroes/

Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.

Example 1:
	Input: nums = [0,1,0,3,12]
	Output: [1,3,12,0,0]

Example 2:
	Input: nums = [0]
	Output: [0]

Constraints:
 * 1 <= nums.length <= 10^4
 * -2^31 <= nums[i] <= 2^31 - 1

Follow up: Could you minimize the total number of operations done?
*/
func main() {
	cases := [][]int{
		{0, 1, 0, 3, 12},
		{0, 0, 1},
	}
	for _, nums := range cases {
		moveZeroes3(nums)
		fmt.Println(nums)
	}
}

// based on bubble sort
// swap nums[j], nums[j+1] when nums[j] is 0
// O(N^2) O(1)
func moveZeroes(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] == 0 {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// refer: https://leetcode.com/problems/move-zeroes/solution/
func moveZeroes2(nums []int) {
	lastNonZeroFoundAt := 0
	// If the current element is not 0, then we need to
	// append it just in front of last non 0 element we found.
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastNonZeroFoundAt] = nums[i]
			lastNonZeroFoundAt++
		}
	}
	// After we have finished processing new elements,
	// all the non-zero elements are already at beginning of array.
	// We just need to fill remaining array with 0's.
	for i := lastNonZeroFoundAt; i < len(nums); i++ {
		nums[i] = 0
	}
}

// double pointer
// 左指针指向当前已经处理好的序列的尾部，右指针指向待处理序列的头部
// 右指针不断向右移动，每次右指针指向非零数，则将左右指针对应的数交换，同时左指针右移。
// 每次交换，都是将左指针的零与右指针的非零数交换，且非零数的相对顺序并未改变。有如下性质：
// 	左指针左边均为非零数；
// 	右指针左边直到左指针处均为零
// O(N) O(1)
// refer: https://leetcode.cn/problems/move-zeroes/solutions/489622/yi-dong-ling-by-leetcode-solution/
func moveZeroes3(nums []int) {
	left, right := 0, 0

	for ; right < len(nums); right++ {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
}
