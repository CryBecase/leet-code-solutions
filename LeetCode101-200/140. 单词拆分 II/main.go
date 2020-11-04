package main

import "strings"

// https://leetcode-cn.com/problems/word-break-ii/

func wordBreak(s string, wordDict []string) (sentences []string) {
	helper := map[string]struct{}{}
	for _, w := range wordDict {
		helper[w] = struct{}{}
	}

	n := len(s)
	dp := make([][][]string, n)

	backtrack := func(index int) [][]string { return nil }
	backtrack = func(index int) [][]string {
		if dp[index] != nil {
			return dp[index]
		}
		wordsList := [][]string{}
		for i := index + 1; i < n; i++ {
			word := s[index:i]
			if _, has := helper[word]; has {
				for _, nextWords := range backtrack(i) {
					wordsList = append(wordsList, append([]string{word}, nextWords...))
				}
			}
		}
		word := s[index:]
		if _, has := helper[word]; has {
			wordsList = append(wordsList, []string{word})
		}
		dp[index] = wordsList
		return wordsList
	}
	for _, words := range backtrack(0) {
		sentences = append(sentences, strings.Join(words, " "))
	}
	return
}

func main() {

}
