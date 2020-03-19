package main

import "fmt"

// https://leetcode-cn.com/problems/longest-palindrome/

func longestPalindrome(s string) int {
	helper := make(map[int32]int)
	for _, v := range s {
		helper[v]++
	}
	result := 0
	canAddOne := true
	for _, v := range helper {
		if canAddOne && v%2 == 1 {
			canAddOne = false
			result++
		}
		result = result + v/2*2
	}
	return result
}

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
}
