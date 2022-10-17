/**
https://leetcode.com/problems/group-anagrams/

Given an array of strings strs, group the anagrams together. You can return the answer in any order.
An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
    typically using all the original letters exactly once.

Example 1:
	Input: strs = ["eat","tea","tan","ate","nat","bat"]
	Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Example 2:
	Input: strs = [""]
	Output: [[""]]

Example 3:
	Input: strs = ["a"]
	Output: [["a"]]

Constraints:
 * 1 <= strs.length <= 10^4
 * 0 <= strs[i].length <= 100
 * strs[i] consists of lowercase English letters.
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	cases := [][]string{
		{"eat", "tea", "tan", "ate", "nat", "bat"},
		{""},
		{"a"},
	}
	for _, strs := range cases {
		fmt.Println(strs)
		fmt.Printf("\t%v\n", groupAnagrams(strs))
	}
}

// hash map
// note: way of gen key
// O(n(k+|Σ|)) O(n(k+|Σ|))
func groupAnagrams(strs []string) [][]string {
	strMap := make(map[[26]int][]string)
	for _, str := range strs {
		key := [26]int{}
		for _, s := range str {
			key[s-'a']++
		}
		strMap[key] = append(strMap[key], str)
	}

	result := make([][]string, 0, len(strMap))
	for _, strs := range strMap {
		result = append(result, strs)
	}
	return result
}

// other ways
// sort str and the sorted str as a key
// O(nklogk) O(nk)
func groupAnagrams2(strs []string) [][]string {
	mp := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
