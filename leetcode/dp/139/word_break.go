/**
https://leetcode.com/problems/word-break/

Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.

Example 1:
	Input: s = "leetcode", wordDict = ["leet","code"]
	Output: true
	Explanation: Return true because "leetcode" can be segmented as "leet code".

Example 2:
	Input: s = "applepenapple", wordDict = ["apple","pen"]
	Output: true
	Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
	Note that you are allowed to reuse a dictionary word.

Example 3:
	Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
	Output: false

Constraints:
 * 1 <= s.length <= 300
 * 1 <= wordDict.length <= 1000
 * 1 <= wordDict[i].length <= 20
 * s and wordDict[i] consist of only lowercase English letters.
 * All the strings of wordDict are unique.
*/
package main

import "fmt"

func main() {
	fmt.Println("leetcode", wordBreak3("leetcode", []string{"leet", "code"}))
	fmt.Println("leetcode", wordBreak3("leetcode", []string{"le", "leet", "code"}))
	fmt.Println("applepenapple", wordBreak3("applepenapple", []string{"apple", "pen"}))
	fmt.Println("catsandog", wordBreak3("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}

// dp[i]: s[0:i] could be split to multiple words in the wordDict
// dp[i] = dp[j] && check(s[j:i])
//	 - s[0 .. j .. i-1]
//	 - check(s[j:i]): check s[j:i] is in dict or not
// O(n^2): n is len(s)
// O(n)
func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		wordMap[word] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

// Brute Force
// The naive approach to solve this problem is to use recursion and backtracking. For finding the solution,
// we check every possible prefix of that string in the dictionary of words, if it is found in the dictionary,
// then the recursive function is called for the remaining portion of that string.
// And, if in some function call it is found that the complete string is in dictionary, then it will return true.

// Time complexity: O(2^n): there are n + 1 ways to split it into two parts. At each step, we have a choice:
//	to split or not to split. In the worse case, when all choices are to be checked, that results in O(2^n).
// Space complexity: O(n). The depth of the recursion tree can go up to n.
// refer: https://leetcode.com/problems/word-break/solution/
func wordBreak2(s string, wordDict []string) bool {
	wordMap := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		wordMap[word] = true
	}
	return wordBreakRecur(s, wordMap, 0)
}

func wordBreakRecur(s string, wordDict map[string]bool, start int) bool {
	if start == len(s) {
		return true
	}

	for end := start + 1; end <= len(s); end++ {
		if wordDict[s[start:end]] && wordBreakRecur(s, wordDict, end) {
			return true
		}
	}
	return false
}

// recursion with memoization
// In the previous approach we can see that many subproblems were redundant, i.e we were calling
// the recursive function multiple times for a particular string.
// To avoid this we can use memoization method, where an array memo is used to store the result of the subproblems.
// Now, when the func is called again for a particular string, value will be fetched and returned using the memo map,
// if its value has been already evaluated.

// Time complexity: O(n^3). Size of recursion tree can go up to n^2
// Space complexity: O(n). The depth of recursion tree can go up to nn.
func wordBreak3(s string, wordDict []string) bool {
	wordMap := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		wordMap[word] = true
	}

	return wordBreakMemo(s, wordMap, 0, make(map[int]bool, len(s)))
}

func wordBreakMemo(s string, wordDict map[string]bool, start int, memo map[int]bool) bool {
	if start == len(s) {
		return true
	}
	if v, ok := memo[start]; ok {
		return v
	}

	for end := start + 1; end <= len(s); end++ {
		if wordDict[s[start:end]] && wordBreakMemo(s, wordDict, end, memo) {
			memo[start] = true
			return true
		}
	}

	memo[start] = false
	return false
}
