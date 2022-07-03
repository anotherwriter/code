package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.com/problems/two-sum/
https://leetcode.cn/problems/two-sum/

Next challenges:
	3Sum, 4Sum

Desc:
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example 1:
	Input: nums = [2,7,11,15], target = 9
	Output: [0,1]
	Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:
	Input: nums = [3,2,4], target = 6
	Output: [1,2]

Example 3:
	Input: nums = [3,3], target = 6
	Output: [0,1]

Constraints:
 * 2 <= nums.length <= 10^4
 * -10^9 <= nums[i] <= 10^9
 * -10^9 <= target <= 10^9
 * Only one valid answer exists.

Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?

*/

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum4(nums, target))

	nums = []int{3, 2, 4}
	target = 6
	fmt.Println(twoSum4(nums, target))

	nums = []int{3, 3}
	target = 6
	fmt.Println(twoSum4(nums, target))
}

// based on search
// Time Complexity: O(N^2)
// Extra Space: O(1)
func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		j := search(nums, i+1, target-v)
		if j != -1 {
			return []int{i, j}
		}
	}

	return nil
}

func search(nums []int, startIndex int, target int) int {
	for i := startIndex; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}

// sort + double pointers
// 10ms
// Time Complexity: O(N)
// Extra Space: O(N)
func twoSum2(nums []int, target int) []int {
	type Elem struct {
		Index int
		Value int
	}
	elems := make([]*Elem, len(nums))
	for i, v := range nums {
		elems[i] = &Elem{
			Index: i,
			Value: v,
		}
	}
	sort.Slice(elems, func(i, j int) bool {
		if elems[i].Value < elems[j].Value {
			return true
		}
		return false
	})

	for i, j := 0, len(elems)-1; i < j; {
		elem1 := elems[i]
		elem2 := elems[j]

		sum := elem1.Value + elem2.Value
		if sum > target {
			j--
		} else if sum < target {
			i++
		} else {
			return []int{elem1.Index, elem2.Index}
		}
	}

	return nil
}

// optimize twoSum2
// 13ms, spend more runtime than twoSum2...
func twoSum3(nums []int, target int) []int {
	indexMap := make(map[int]int, len(nums)) // value => index
	for i, v := range nums {
		// since only one valid answer exists
		indexMap[v] = i
	}

	for i, v := range nums {
		j, ok := indexMap[target-v]
		if ok && j != i {
			return []int{i, j}
		}
	}

	return nil
}

// optimize twoSum4 by storing target - nums[i] => index
// 8ms
func twoSum4(nums []int, target int) []int {
	indexMap := make(map[int]int, len(nums)) // target - nums[i] => index

	for i, v := range nums {
		left, ok := indexMap[v]
		if !ok {
			indexMap[target-v] = i
		} else {
			return []int{left, i}
		}
	}

	return nil
}
