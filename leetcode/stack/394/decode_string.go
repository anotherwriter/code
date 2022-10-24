/**
https://leetcode.com/problems/decode-string/

Given an encoded string, return its decoded string.

The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times.
Note that k is guaranteed to be a positive integer.

You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed,
	etc. Furthermore, you may assume that the original data does not contain any digits and that digits are only for those
	repeat numbers, k. For example, there will not be input like 3a or 2[4].

The test cases are generated so that the length of the output will never exceed 105.

Example 1:
	Input: s = "3[a]2[bc]"
	Output: "aaabcbc"

Example 2:
	Input: s = "3[a2[c]]"
	Output: "accaccacc"

Example 3:
	Input: s = "2[abc]3[cd]ef"
	Output: "abcabccdcdcdef"

Constraints:
 * 1 <= s.length <= 30
 * s consists of lowercase English letters, digits, and square brackets '[]'.
 * s is guaranteed to be a valid input.
 * All the integers in s are in the range [1, 300].
*/
package main

import "fmt"

func main() {
	fmt.Println(decodeString("3[a]2[bc]"))
	fmt.Println(decodeString("3[a2[c]]"))
	fmt.Println(decodeString("2[abc]3[cd]ef"))
}

func decodeString(s string) string {
	var result string
	var multiStack []int
	var ansStack []string

	var multi int
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			multi = multi*10 + int(ch-'0')
		} else if ch == '[' {
			ansStack = append(ansStack, result)
			multiStack = append(multiStack, multi)
			result = ""
			multi = 0
		} else if ch >= 'a' && ch <= 'z' {
			result += fmt.Sprintf("%c", ch)
		} else if ch == ']' {
			var ansTmp string
			if len(ansStack) > 0 {
				ansTmp = ansStack[len(ansStack)-1]
				ansStack = ansStack[:len(ansStack)-1]
			}

			var multiTmp int
			if len(multiStack) > 0 {
				multiTmp = multiStack[len(multiStack)-1]
				multiStack = multiStack[:len(multiStack)-1]
			}

			for i := 0; i < multiTmp; i++ {
				ansTmp += result
			}
			result = ansTmp
		}
	}

	return result
}
