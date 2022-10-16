package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.com/problems/3sum/
https://leetcode.cn/problems/3sum/

Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]]
such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.

Example 1:
	Input: nums = [-1,0,1,2,-1,-4]
	Output: [[-1,-1,2],[-1,0,1]]

Example 2:
	Input: nums = []
	Output: []

Example 3:
	Input: nums = [0]
	Output: []

Constraints:
 * 0 <= nums.length <= 3000
 * -10^5 <= nums[i] <= 10^5
*/
func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum4(nums))
}

// O(N^3)
func threeSum1(nums []int) [][]int {
	sort.Ints(nums)

	var result [][]int
	var i, j, k int
	for ; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] {
			for j = i + 1; j < len(nums); j++ {
				if j == i+1 || nums[j] != nums[j-1] {
					for k = j + 1; k < len(nums); k++ {
						if k == j+1 || nums[k] != nums[k-1] {
							if nums[i]+nums[j]+nums[k] == 0 {
								result = append(result, []int{nums[i], nums[j], nums[k]})
								break // a + b + c = 0. a, b are determined => c will be determined.
							}
						}
					}
				}
			}
		}
	}

	return result
}

// based on map
func threeSum2(nums []int) [][]int {
	num2index := make(map[int]int, len(nums)) // num => index
	for i, v := range nums {
		num2index[v] = i
	}

	resultMap := map[string][]int{} // v1,v2,v3 => []int{v1, v2, v3}
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			target := -v - nums[j]
			if k, ok := num2index[target]; ok && k > j {
				tmp := []int{v, nums[j], target}
				sort.Ints(tmp)
				resultMap[fmt.Sprintf("%d,%d,%d", tmp[0], tmp[1], tmp[2])] = tmp
			}
		}
	}

	result := make([][]int, 0, len(resultMap))
	for _, v := range resultMap {
		result = append(result, v)
	}
	return result
}

// base sort + two pointers. use map to implement unique triplets
// less runtime than threeSum1
func threeSum3(nums []int) [][]int {
	sort.Ints(nums)

	resultMap := make(map[string][]int)

	var j, k, sum int
	for i, v := range nums {
		j, k = i+1, len(nums)-1
		for j < k {
			sum = nums[j] + nums[k]
			if sum < -v {
				j++
			} else if sum > -v {
				k--
			} else {
				resultMap[fmt.Sprintf("%d,%d,%d", v, nums[j], nums[k])] = []int{v, nums[j], nums[k]}
				// we can do that since the question need unique triplets
				j++
				k--
			}
		}
	}

	result := make([][]int, 0, len(resultMap))
	for _, v := range resultMap {
		result = append(result, v)
	}
	return result
}

// optimize threeSum2
// base sort + two pointers without using map
func threeSum4(nums []int) [][]int {
	sort.Ints(nums)

	var result [][]int
	var i, j, k, sum int
	for ; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k = i+1, len(nums)-1
		for j < k {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			sum = nums[j] + nums[k]
			if sum < -nums[i] {
				j++
			} else if sum > -nums[i] {
				k--
			} else {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			}
		}
	}
	return result
}
