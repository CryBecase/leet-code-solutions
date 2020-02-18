package main

import (
	"fmt"
)

/**
https://leetcode-cn.com/problems/longest-word-in-dictionary-through-deleting
给定一个字符串和一个字符串字典，找到字典里面最长的字符串，该字符串可以通过删除给定字符串的某些字符来得到。如果答案不止一个，返回长度最长且字典顺序最小的字符串。如果答案不存在，则返回空字符串。

示例 1:

输入:
s = "abpcplea", d = ["ale","apple","monkey","plea"]

输出:
"apple"
示例 2:

输入:
s = "abpcplea", d = ["a","b","c"]

输出:
"a"
说明:

所有输入的字符串只包含小写字母。
字典的大小不会超过 1000。
所有输入的字符串长度不会超过 1000。
*/

func findLongestWord(s string, d []string) string {
	if s == "" || len(d) == 0 {
		return ""
	}

	result := ""
	for i := range d {
		p1, p2 := 0, 0
		if len(d[i]) < len(result) {
			continue
		}
		for p1 < len(s) && p2 < len(d[i]) {
			if s[p1] == d[i][p2] {
				p2++
			}
			p1++
		}
		if p2 == len(d[i]) {
			// 可以通过删除s中的字母得到 d[i]
			if len(d[i]) > len(result) {
				result = d[i]
			} else if len(d[i]) == len(result) {
				if d[i] < result {
					result = d[i]
				}
			}
		}
	}
	return result
}

func main() {
	s := "abpcplea"
	d := []string{"ba", "ab", "monkey", "plea"}
	fmt.Println(findLongestWord(s, d))
}
