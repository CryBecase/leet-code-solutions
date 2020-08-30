package main

import "strings"

// https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/

func reverseWords(s string) string {
	word := make([]byte, 0)
	builder := strings.Builder{}
	for i := range s {
		if s[i] == ' ' {
			for j := len(word) - 1; j >= 0; j-- {
				builder.WriteByte(word[j])
			}
			builder.WriteByte(' ')
			word = word[0:0]
			continue
		}

		word = append(word, s[i])
	}
	for j := len(word) - 1; j >= 0; j-- {
		builder.WriteByte(word[j])
	}
	return builder.String()
}

func main() {

}
