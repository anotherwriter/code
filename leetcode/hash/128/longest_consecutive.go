/**
https://leetcode.com/problems/longest-consecutive-sequence/

Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
You must write an algorithm that runs in O(n) time.

Example 1:
	Input: nums = [100,4,200,1,3,2]
	Output: 4
	Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.

Example 2:
	Input: nums = [0,3,7,2,5,8,4,6,0,1]
	Output: 9


Constraints:
 * 0 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
*/

package main

import "fmt"

func main() {
	cases := [][]int{
		{100, 4, 200, 1, 3, 2},
		{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
	}
	for _, nums := range cases {
		fmt.Println(longestConsecutive1(nums))
	}
}

// O(n^2)
func longestConsecutive(nums []int) int {
	numMap := make(map[int]bool, len(nums))
	for _, num := range nums {
		numMap[num] = true
	}

	var maxLen int
	for _, num := range nums {
		j := num + 1
		for ; numMap[j]; j++ {
		}
		if j-num > maxLen {
			maxLen = j - num
		}
	}
	return maxLen
}

// O(n)
func longestConsecutive1(nums []int) int {
	numMap := make(map[int]bool, len(nums))
	for _, num := range nums {
		numMap[num] = true
	}

	var maxLen int
	for _, num := range nums {
		if numMap[num-1] {
			continue
		}

		j := num + 1
		for ; numMap[j]; j++ {
		}
		if j-num > maxLen {
			maxLen = j - num
		}
	}
	return maxLen
}
