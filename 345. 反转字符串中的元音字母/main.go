package main

import "fmt"

/**
https://leetcode-cn.com/problems/reverse-vowels-of-a-string
编写一个函数，以字符串作为输入，反转该字符串中的元音字母。

示例 1:

输入: "hello"
输出: "holle"
示例 2:

输入: "leetcode"
输出: "leotcede"
说明:
元音字母不包含字母"y"。
*/
// a e i o u A E I O U
func reverseVowels(s string) string {
	if len(s) == 0 {
		return s
	}
	n := len(s) - 1
	i := 0
	ss := []byte(s)
	for i < n {
		if isaeiou(s[i]) && isaeiou(s[n]) {
			ss[i], ss[n] = ss[n], ss[i]
			i++
			n--
		} else if !isaeiou(s[i]) {
			i++
		} else if !isaeiou(s[n]) {
			n--
		}
	}
	return string(ss)
}

func isaeiou(s byte) bool {
	if s == 'a' || s == 'A' || s == 'e' || s == 'E' || s == 'i' || s == 'I' || s == 'o' || s == 'O' || s == 'u' || s == 'U' {
		return true
	}
	return false
}

func main() {
	fmt.Println(reverseVowels("leetcode"))
}
