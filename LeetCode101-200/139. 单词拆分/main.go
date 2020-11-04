package main

// https://leetcode-cn.com/problems/word-break/

func wordBreak(s string, wordDict []string) bool {
	helper := make(map[string]bool)
	for _, word := range wordDict {
		helper[word] = true
	}
	// dp[i] 表示从 前i个字符组成的字符串能否被分割成多个字典中出现的单词
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && helper[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func main() {

}
