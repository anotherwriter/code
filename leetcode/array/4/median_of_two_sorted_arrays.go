package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.com/problems/median-of-two-sorted-arrays/
https://leetcode.cn/problems/median-of-two-sorted-arrays/

Next challenges:


Desc:
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).

Example 1:
	Input: nums1 = [1,3], nums2 = [2]
	Output: 2.00000
	Explanation: merged array = [1,2,3] and median is 2.

Example 2:
	Input: nums1 = [1,2], nums2 = [3,4]
	Output: 2.50000
	Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

Constraints:
 * nums1.length == m
 * nums2.length == n
 * 0 <= m <= 1000
 * 0 <= n <= 1000
 * 1 <= m + n <= 2000
 * -106 <= nums1[i], nums2[i] <= 106

*/

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays2(nums1, nums2))

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	fmt.Println(findMedianSortedArrays2(nums1, nums2))
}

//#region findMedianSortedArrays1
// based on merging two sorted arrays
// Time Complexity: O(m+n)
// Extra Space: O(m+n). TODO: optimize to O(1)
// leetcode: 39ms, 5.8MB
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	merged := mergeSortedArrays(nums1, nums2)

	n := len(merged)
	if n%2 == 0 {
		return float64(merged[(n-1)/2]+merged[n/2]) / 2.0
	}
	return float64(merged[n/2])
}

func mergeSortedArrays(nums1 []int, nums2 []int) []int {
	merged := make([]int, len(nums1)+len(nums2))

	var i, j int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			merged[i+j] = nums1[i]
			i++
		} else {
			merged[i+j] = nums2[j]
			j++
		}
	}
	for ; i < len(nums1); i++ {
		merged[i+j] = nums1[i]
	}
	for ; j < len(nums2); j++ {
		merged[i+j] = nums2[j]
	}

	return merged
}

//#endregion

// 37ms, 5.3MB
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	var flagIndex int

	size := len(nums1) + len(nums2)
	if size%2 == 0 {
		flagIndex = (size - 1) / 2
	} else {
		flagIndex = size / 2
	}

	var i, j int
	for i < len(nums1) && j < len(nums2) && i+j < flagIndex {
		if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	for i < len(nums1) && i+j < flagIndex {
		i++
	}
	for j < len(nums2) && i+j < flagIndex {
		j++
	}

	// i + j must be flagIndex
	if size%2 == 0 {
		values := make([]int, 0, 4)
		if i < len(nums1) {
			values = append(values, nums1[i])
		}
		if i+1 < len(nums1) {
			values = append(values, nums1[i+1])
		}
		if j < len(nums2) {
			values = append(values, nums2[j])
		}
		if j+1 < len(nums2) {
			values = append(values, nums2[j+1])
		}
		sort.Ints(values)
		return float64(values[0]+values[1]) / 2.0
	}

	if i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			return float64(nums1[i])
		}
		return float64(nums2[j])
	}
	if i < len(nums1) {
		return float64(nums1[i])
	}
	return float64(nums2[j])
}

// binary search
func findMedianSortedArrays3(nums1 []int, nums2 []int) float64 {
	return -1
}
