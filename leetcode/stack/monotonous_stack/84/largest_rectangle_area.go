/**
https://leetcode.com/problems/largest-rectangle-in-histogram/

Given an array of integers heights representing the histogram's bar height where the width of each bar is 1, return the area of the largest rectangle in the histogram.

Example 1:
	Input: heights = [2,1,5,6,2,3]
	Output: 10
	Explanation: The above is a histogram where width of each bar is 1.
	The largest rectangle is shown in the red area, which has an area = 10 units.

Example2:
	Input: heights = [2,4]
	Output: 4

Constraints:
 * 1 <= heights.length <= 105
 * 0 <= heights[i] <= 104
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	cases := [][]int{
		{2, 1, 5, 6, 2, 3}, // 10
		{0, 1, 0, 1},       // 1
		{2, 4},             // 4
	}

	for _, heights := range cases {
		fmt.Println(largestRectangleArea2(heights))
	}
}

// O(n^2)
// 枚举宽
func largestRectangleArea(heights []int) int {
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		minHeight := heights[i]
		for j := i; j < len(heights); j++ {
			minHeight = int(math.Min(float64(minHeight), float64(heights[j])))
			area := minHeight * (j - i + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

// O(n^2)
// 枚举高
// 使用一重循环枚举某一根柱子，将其固定为矩形的高度 h
// 随后我们从这跟柱子开始向两侧延伸，直到遇到高度小于 h 的柱子，就确定了矩形的左右边界
// 如果左右边界之间的宽度为 w，那么对应的面积为 w * h
// refer: https://leetcode.cn/problems/largest-rectangle-in-histogram/solution/zhu-zhuang-tu-zhong-zui-da-de-ju-xing-by-leetcode-/
func largestRectangleArea1(heights []int) int {
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		height := heights[i]

		left, right := i, i
		for left-1 >= 0 && heights[left-1] >= height {
			left--
		}
		for right+1 < len(heights) && heights[right+1] >= height {
			right++
		}

		area := height * (right - left + 1)
		maxArea = int(math.Max(float64(maxArea), float64(area)))
	}

	return maxArea
}

// optimize largestRectangleArea1 based on monotonous increasing stack
// 通过单调栈的思维找到栈顶元素左边第一个比它矮的位置，和右边比它矮的位置，然后我们依次来遍历这个柱状图中的高度数组，
// 我们就能求出以每个元素做为高度的矩形的面积，然后选出来一个最大的，就能求解了
// O(n) O(n)
func largestRectangleArea2(heights []int) int {
	lefts, rights := make([]int, len(heights)), make([]int, len(heights)) // save indexes

	// init lefts
	var monoStack []int
	for i, height := range heights {
		for len(monoStack) > 0 && height <= heights[monoStack[len(monoStack)-1]] {
			// pop stack if current height <= top value
			monoStack = monoStack[:len(monoStack)-1]
		}

		if len(monoStack) == 0 {
			lefts[i] = -1
		} else {
			lefts[i] = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}

	// init rights
	monoStack = []int{}
	for i := len(heights) - 1; i >= 0; i-- {
		for len(monoStack) > 0 && heights[i] <= heights[monoStack[len(monoStack)-1]] {
			monoStack = monoStack[:len(monoStack)-1]
		}

		if len(monoStack) == 0 {
			rights[i] = len(heights)
		} else {
			rights[i] = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}

	maxArea := 0
	for i, height := range heights {
		maxArea = max(maxArea, (rights[i]-lefts[i]-1)*height)
	}
	return maxArea
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
