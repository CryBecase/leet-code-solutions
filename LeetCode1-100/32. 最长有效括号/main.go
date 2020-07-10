package main

import "fmt"

// https://leetcode-cn.com/problems/longest-valid-parentheses/

func longestValidParentheses(s string) int {
	// dp[i] 表示以第i个符号为子串结尾的最长有效括号长度
	// dp[i] = 2 + dp[i-1] + dp[i-dp[i-1]-2]
	dp := make([]int, len(s))
	max := 0
	for i, char := range s {
		if char == ')' && i-1 >= 0 && i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
			dp[i] = 2 + dp[i-1]
			if i-dp[i-1]-2 >= 0 {
				dp[i] += dp[i-dp[i-1]-2]
			}
			if dp[i] > max {
				max = dp[i]
			}
		}
	}
	return max
}

func main() {
	fmt.Println(longestValidParentheses(")()())"))
}
