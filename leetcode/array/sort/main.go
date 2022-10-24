package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{0, 1, 0, 3, 15, 12}
	//quickSort(nums, 0, len(nums)-1)
	mergeSort(nums, 0, len(nums)-1)
	//swapSort(nums)
	fmt.Println(nums)
}

func insertSort(nums []int) {
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] >= nums[i-1] {
			continue
		}

		// insert nums[i] to nums[0, i-1]
		var j int
		tmpValue := nums[i]
		for j = i - 1; j >= 0 && nums[j] > tmpValue; j-- {
			nums[j+1] = nums[j]
		}
		nums[j+1] = tmpValue
	}
}

func swapSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ { // nums - 1 times sort
		for j := i + 1; j < len(nums); j++ { // nums[i] is minimum after comparing nums[i] with nums[i+1,...]
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

func bubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ { // nums - 1 times sort
		for j := 0; j < len(nums)-i-1; j++ { // nums[len-i-1] is maximum after comparing
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func quickSort(nums []int, p, r int) {
	if p < r {
		q := partition(nums, p, r)

		quickSort(nums, p, q-1)
		quickSort(nums, q+1, r)
	}
}

func partition(nums []int, p, r int) int {
	x := nums[r] // x为划分元

	j := p - 1 // nums[p, ...j-1] <= x
	for i := p; i < r; i++ {
		if nums[i] < x {
			j++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[j+1], nums[r] = nums[r], nums[j+1]

	return j + 1
}

func mergeSort(nums []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		mergeSort(nums, p, q)
		mergeSort(nums, q+1, r)
		mergeWithGuard(nums, p, q, r)
	}
}

// nums[p...q] nums[q+1...r] is sorted, merge them and make nums[p...r] is sorted
func merge(nums []int, p, q, r int) {
	nums1 := make([]int, q-p+1)
	nums2 := make([]int, r-q)

	for i := p; i <= q; i++ {
		nums1[i-p] = nums[i]
	}
	for i := q + 1; i <= r; i++ {
		nums2[i-q-1] = nums[i]
	}

	var i, j int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			nums[i+j+p] = nums1[i]
			i++
		} else {
			nums[i+j+p] = nums2[j]
			j++
		}
	}
	for ; i < len(nums1); i++ {
		nums[i+j+p] = nums1[i]
	}
	for ; j < len(nums2); j++ {
		nums[i+j+p] = nums2[j]
	}
}

// nums[p...q] nums[q+1...r] is sorted, merge them and make nums[p...r] is sorted
func mergeWithGuard(nums []int, p, q, r int) {
	n1, n2 := q-p+1, r-q
	nums1 := make([]int, n1+1)
	nums2 := make([]int, n2+1)

	for i := p; i <= q; i++ {
		nums1[i-p] = nums[i]
	}
	for i := q + 1; i <= r; i++ {
		nums2[i-q-1] = nums[i]
	}
	nums1[n1], nums2[n2] = math.MaxInt, math.MaxInt

	var i, j int
	for k := p; k <= r; k++ {
		if nums1[i] <= nums2[j] {
			nums[k] = nums1[i]
			i++
		} else {
			nums[k] = nums2[j]
			j++
		}
	}
}
