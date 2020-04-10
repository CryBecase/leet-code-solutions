package main

/**
https://leetcode-cn.com/problems/longest-palindromic-substring/

给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
*/

/**
扩散法
注意边界
*/
func longestPalindrome(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1, len2 := expandAroundCenter(s, i, i), expandAroundCenter(s, i, i+1)
		maxLen := max(len1, len2)
		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	L, R := left, right
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}
	return R - L - 1
}

func max(n, m int) int {
	if n >= m {
		return n
	} else {
		return m
	}
}

func main() {

}
