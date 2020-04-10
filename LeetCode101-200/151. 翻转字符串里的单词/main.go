package main

import (
	"fmt"
	"strings"
)

// https://leetcode-cn.com/problems/reverse-words-in-a-string/

func reverseWords(s string) string {
	arr := strings.Fields(s)
	b := strings.Builder{}
	for i := len(arr) - 1; i >= 0; i-- {
		b.WriteString(arr[i])
		if i != 0 {
			b.WriteString(" ")
		}
	}
	return b.String()
}

func main() {
	words := reverseWords("  hello world! ")
	fmt.Println(words)
}
