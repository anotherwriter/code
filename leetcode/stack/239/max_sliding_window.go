/**
https://leetcode.com/problems/sliding-window-maximum/

  You are given an array of integers nums, there is a sliding window of size k
which is moving from the very left of the array to the very right. You can only see the k numbers in the window.
Each time the sliding window moves right by one position.

Return the max sliding window.

Example 1:
	Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
	Output: [3,3,5,5,6,7]
	Explanation:
		Window position                Max
		---------------               -----
		[1  3  -1] -3  5  3  6  7       3
		 1 [3  -1  -3] 5  3  6  7       3
		 1  3 [-1  -3  5] 3  6  7       5
		 1  3  -1 [-3  5  3] 6  7       5
		 1  3  -1  -3 [5  3  6] 7       6
		 1  3  -1  -3  5 [3  6  7]      7
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow([]int{7, 2, 4}, 2))
}

// based on sort
func maxSlidingWindow(nums []int, k int) []int {
	window := make([]int, k)

	result := make([]int, 0, len(nums)-k+1)
	for i := 0; i <= len(nums)-k; i++ {
		for j := i; j < i+k; j++ {
			window[j-i] = nums[j]
		}
		sort.Ints(window)
		result = append(result, window[k-1])
	}

	return result
}

// based on queue
// https://leetcode.cn/problems/sliding-window-maximum/solutions/543426/hua-dong-chuang-kou-zui-da-zhi-by-leetco-ki6m/
func maxSlidingWindow2(nums []int, k int) []int {
	indexQueue := make([]int, 0, k)
	push := func(i int) {
		for len(indexQueue) > 0 && nums[i] >= nums[indexQueue[len(indexQueue)-1]] {
			// pop
			indexQueue = indexQueue[:len(indexQueue)-1]
		}
		indexQueue = append(indexQueue, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	result := make([]int, 1, n-k+1)
	result[0] = nums[indexQueue[0]]
	for i := k; i < n; i++ {
		push(i)
		for indexQueue[0] <= i-k {
			indexQueue = indexQueue[1:]
		}
		result = append(result, nums[indexQueue[0]])
	}

	return result
}
