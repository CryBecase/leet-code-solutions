package main

/**
https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

// 滑动窗口
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	result := 0
	window := make(map[byte]int)
	for i, j := 0, 0; j < len(s); j++ {
		// v 是重复的数的下标
		if v, ok := window[s[j]]; ok {
			// 现在经过的 byte 可能和之前走过的 byte 是一样的，不能让 i 返回去
			if i < v+1 {
				i = v + 1
			}
		}
		if result < j-i+1 {
			result = j - i + 1
		}
		window[s[j]] = j
	}

	return result
}

func main() {

}
