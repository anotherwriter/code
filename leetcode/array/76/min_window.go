package main

import (
	"fmt"
	"math"
)

/**
https://leetcode.com/problems/minimum-window-substring/
https://leetcode.cn/problems/minimum-window-substring/

Given two strings s and t of lengths m and n respectively, return the minimum window substring of s such that
  every character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".

The testcases will be generated such that the answer is unique.

A substring is a contiguous sequence of characters within the string.

Example 1:
	Input: s = "ADOBECODEBANC", t = "ABC"
	Output: "BANC"
	Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.

Example 2:
	Input: s = "a", t = "a"
	Output: "a"
	Explanation: The entire string s is the minimum window.

Example 3:
	Input: s = "a", t = "aa"
	Output: ""
	Explanation: Both 'a's from t must be included in the window.
	Since the largest window of s only has one 'a', return empty string.

Constraints:
 * m == s.length
 * n == t.length
 * 1 <= m, n <= 10^5
 * s and t consist of uppercase and lowercase English letters.

Follow up: Could you find an algorithm that runs in O(m + n) time?
*/
func main() {
	type Case struct {
		s string
		t string
	}
	cases := []Case{
		{s: "ADOBECODEBANC", t: "ABC"},
		{s: "a", t: "a"},
		{s: "a", t: "aa"},
	}

	for _, c := range cases {
		fmt.Println(minWindow(c.s, c.t))
	}
}

func minWindow(s, t string) string {
	return minWindow4(s, t)
}

// O(n^2)
func minWindow1(s, t string) string {
	if len(t) > len(s) {
		return ""
	}
	if !isContain(s, t) {
		return ""
	}

	result := s
	strs := allSubStrings(s)
	for str := range strs {
		if !isContain(str, t) {
			continue
		}

		if len(result) > len(str) {
			result = str
		}
	}

	return result
}

func allSubStrings(s string) map[string]struct{} {
	result := map[string]struct{}{}
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			result[s[i:j]] = struct{}{}
		}
	}

	return result
}

func isContain(s, t string) bool {
	strMap := make(map[uint8]int, len(s))
	for i := 0; i < len(s); i++ {
		strMap[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		cnt, ok := strMap[t[i]]
		if !ok {
			return false
		}

		cnt--
		if cnt == 0 {
			delete(strMap, t[i])
			continue
		}
		strMap[t[i]] = cnt
	}

	return true
}

// sliding window
// Time Limit Exceeded
func minWindow2(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	targetMap := make(map[uint8]int, len(t)) // target char => count
	for i := 0; i < len(t); i++ {
		targetMap[t[i]]++
	}

	result := "***"
	for i, j := 0, 0; i < len(s)-len(t)+1 && j < len(s); {
		currStr := s[i : j+1]
		if isContainTarget(currStr, targetMap) {
			if result == "***" {
				result = currStr
			} else if len(currStr) < len(result) {
				result = currStr
			}
			i++
			continue
		}
		j++
	}

	if result == "***" {
		return ""
	}
	return result
}

func isContainTarget(s string, targetMap map[uint8]int) bool {
	strMap := make(map[uint8]int, len(s))
	for i := 0; i < len(s); i++ {
		strMap[s[i]]++
	}

	for k, v := range targetMap {
		cnt, ok := strMap[k]
		if !ok || cnt < v {
			return false
		}
	}
	return true
}

// sliding window: optimize check function
// 337ms 3.9MB
func minWindow3(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	strMap := make(map[uint8]int, 128)    // target char => count
	targetMap := make(map[uint8]int, 128) // target char => count
	for i := 0; i < len(t); i++ {
		targetMap[t[i]]++
	}

	check := func() bool {
		if len(targetMap) > len(strMap) {
			return false
		}

		for k, v := range targetMap {
			cnt, ok := strMap[k]
			if !ok || cnt < v {
				return false
			}
		}
		return true
	}

	result := "***"
	i, j := 0, 0
	for ; j < len(s); j++ {
		strMap[s[j]]++
		if !check() {
			continue
		}

		tmp := s[i : j+1]
		if result == "***" {
			result = tmp
		} else if len(tmp) < len(result) {
			result = tmp
		}

		for i < j {
			strMap[s[i]]--
			i++

			if !check() {
				break
			}
			tmp := s[i : j+1]
			if len(tmp) < len(result) {
				result = tmp
			}
			continue
		}
	}

	if result == "***" {
		return ""
	}
	return result
}

// 127ms 2.8MB
// refer: https://leetcode.cn/problems/minimum-window-substring/solution/zui-xiao-fu-gai-zi-chuan-by-leetcode-solution/
func minWindow4(s string, t string) string {
	targetMap, strMap := make(map[byte]int, 128), make(map[byte]int, 128)
	for i := 0; i < len(t); i++ {
		targetMap[t[i]]++
	}

	check := func() bool {
		for k, v := range targetMap {
			if strMap[k] < v {
				return false
			}
		}
		return true
	}

	minLen := math.MaxInt32
	ansL, ansR := -1, -1
	for l, r := 0, 0; r < len(s); r++ {
		if r < len(s) && targetMap[s[r]] > 0 {
			strMap[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < minLen {
				minLen = r - l + 1
				ansL, ansR = l, l+minLen
			}
			if _, ok := targetMap[s[l]]; ok {
				strMap[s[l]] -= 1
			}
			l++
		}
	}

	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}

// TODO
func minWindow5(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	strMap := make(map[uint8][]int, len(s))
	for i := 0; i < len(s); i++ {
		strMap[s[i]] = append(strMap[s[i]], i) // char => indexes
	}

	targetMap := make(map[uint8][]int, len(s)) // target char => indexes
	for i := 0; i < len(t); i++ {
		indexes, ok := strMap[t[i]]
		if !ok {
			return ""
		}
		targetMap[t[i]] = indexes
	}

	return ""
}
