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
	sByte := []byte(s)
	vowels := map[byte]interface{}{
		'a': nil,
		'e': nil,
		'i': nil,
		'o': nil,
		'u': nil,
		'A': nil,
		'E': nil,
		'I': nil,
		'O': nil,
		'U': nil,
	}
	L, R := 0, len(sByte)-1
	for L < R {
		for L < R {
			// 找left
			if _, ok := vowels[sByte[L]]; ok {
				// 是元音
				break
			} else {
				// 不是元音
				L++
			}
		}
		for L < R {
			// 找right
			if _, ok := vowels[sByte[R]]; ok {
				break
			} else {
				R--
			}
		}
		sByte[L], sByte[R] = sByte[R], sByte[L]
		L++
		R--
	}
	return string(sByte)
}

func main() {
	fmt.Println(reverseVowels("leetcode"))
}
