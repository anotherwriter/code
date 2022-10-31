/**
https://leetcode.com/problems/subarray-sum-equals-k/

Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.
A subarray is a contiguous non-empty sequence of elements within an array.

Example 1:
	Input: nums = [1,1,1], k = 2
	Output: 2

Example 2:
	Input: nums = [1,2,3], k = 3
	Output: 2

Constraints:
 * 1 <= nums.length <= 2 * 10^4
 * -1000 <= nums[i] <= 1000
 * -10^7 <= k <= 10^7
*/
package main

import "fmt"

func main() {
	cases := [][]interface{}{
		[]interface{}{[]int{1, 1, 1}, 2},  // 2
		[]interface{}{[]int{1, 2, 3}, 3},  // 2
		[]interface{}{[]int{1, -1, 0}, 0}, // 3
		[]interface{}{[]int{0, 0}, 0},     // 3
	}
	for _, data := range cases {
		fmt.Println(subarraySum2(data[0].([]int), data[1].(int)))
	}
}

// O(n^2)
func subarraySum(nums []int, k int) int {
	var cnt int
	for i := 0; i < len(nums); i++ {
		if nums[i] == k {
			cnt++
			// note: don't continue here since the next value may be 0
		}

		sum := nums[i]
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				cnt++
			}
		}
	}
	return cnt
}

// another implement from leetcode
// https://leetcode.cn/problems/subarray-sum-equals-k/solution/he-wei-kde-zi-shu-zu-by-leetcode-solution/
func subarraySum1(nums []int, k int) int {
	count := 0
	for start := 0; start < len(nums); start++ {
		var sum int
		for end := start; end >= 0; end-- {
			sum += nums[end]
			if sum == k {
				count++
			}
		}
	}
	return count
}

// optimize subarraySum1
// O(n) O(n)
// pre[i] = sum[0..i]   pre[i] = pre[i-1] + nums[i]
// if sum[j..i] == k
// 	pre[i] - pre[j-1] == k
// 	pre[j-1] == pre[i] - k
// So that for nums[i], we need to find the count of j that satisfy pre[j-1] == pre[i] - k
func subarraySum2(nums []int, k int) int {
	var count, sum int
	preSumMap := map[int]int{} // sum[0..i] => count
	preSumMap[0] = 1

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if _, ok := preSumMap[sum-k]; ok {
			count += preSumMap[sum-k]
		}
		preSumMap[sum] += 1
	}
	return count
}
