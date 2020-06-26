package main

// https://leetcode-cn.com/problems/find-the-longest-substring-containing-vowels-in-even-counts/

const vowel = "aeiou"

// pre[i][k] 表示在字符串前 i 个字符中，第 k 个元音字母一共出现的次数
func findTheLongestSubstring(s string) int {
	// 创建前缀和数组
	pre := make([][]bool, len(s))
	for i := range pre {
		pre[i] = make([]bool, 5)
	}
	// 初始化前缀和数组
	for i, character := range s {
		if i == 0 {
			if isVowel(character) {
				pre[i][getVowelIndex(character)] = false
			}
		} else {
			if isVowel(character) {
				pre[i][getVowelIndex(character)] = !pre[i-1][getVowelIndex(character)]
			}
		}
	}

}

func getVowel(i int) int32 {
	return int32(vowel[i])
}

func isVowel(c int32) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func getVowelIndex(c int32) int {
	switch c {
	case 'a':
		return 0
	case 'e':
		return 1
	case 'i':
		return 2
	case 'o':
		return 3
	case 'u':
		return 4
	}
	return 0
}

func main() {

}
