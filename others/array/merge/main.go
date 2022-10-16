package main

func main() {

}

// Merge two sorted arrays with O(1) extra space
// https://www.geeksforgeeks.org/merge-two-sorted-arrays-o1-extra-space/

// Merge two sorted arrays
// https://www.geeksforgeeks.org/merge-two-sorted-arrays/

// Method1: O(n1 * n2) Time and O(n1+n2) Extra Space
//	1.Create and array arr3[] of size n1 + n2
//	2.Copy all n1 elements of arr1[] to arr3[]
//	3.Traverse arr2[] and one by one insert elements(like insertion sort) of arr3[] to arr1[]. This step take O(n1 * n2) time

// Method2: O(n1 + n2) Time and O(n1 + n2) Extra Space
//	1.Create and array arr3[] of size n1 + n2
//	2.Simultaneously traverse arr1[] and arr2[]
//		Pick smaller of current elements in arr1[] and arr2[], copy this smaller element to next position in arr3[]
//		and move ahead in arr3[] and the array whose element is picked
//	3.If there are remaining elements in arr1[] or arr2[], copy them also in arr3[]

// Method3: Using Tree Maps, O(nlogn + mlogm) Time and O(N) Extra Space
//	1.Insert elements of both arrays in a map as keys
//	2.Print the keys of the map

func mergeSortedArraysMethod2(nums1 []int, nums2 []int) []int {
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

	//for i, j := 0, 0; i < len(nums1) || j < len(nums2); {
	//	if i == len(nums1) {
	//		merged[i+j] = nums2[j]
	//		j++
	//		continue
	//	}
	//	if j == len(nums2) {
	//		merged[i+j] = nums1[i]
	//		i++
	//		continue
	//	}
	//
	//	if nums1[i] < nums2[j] {
	//		merged[i+j] = nums1[i]
	//		i++
	//	} else {
	//		merged[i+j] = nums2[j]
	//		j++
	//	}
	//}

	return merged
}
