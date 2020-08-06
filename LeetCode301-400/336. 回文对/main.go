package main

// https://leetcode-cn.com/problems/palindrome-pairs/

func palindromePairs(words []string) [][]int {
	// 建立逆序word的索引map
	indices := make(map[string]int)
	for i, word := range words {
		indices[reverse(word)] = i
	}
	result := make([][]int, 0)
	for i, word := range words {
		for j := 0; j <= len(word); j++ {
			if isPal(word[j:len(word)]) {
				if index, ok := indices[word[0:j]]; ok && i != index {
					result = append(result, []int{i, index})
				}
			}
			if j > 0 && isPal(word[0:j]) {
				if index, ok := indices[word[j:len(word)]]; ok && i != index {
					result = append(result, []int{index, i})
				}
			}
		}
	}
	return result
}

func reverse(s string) string {
	n := len(s)
	b := []byte(s)
	for i := 0; i < n/2; i++ {
		b[i], b[n-i-1] = b[n-i-1], b[i]
	}
	return string(b)
}

func isPal(s string) bool {
	if s == "" || len(s) == 1 {
		return true
	}
	L, R := 0, len(s)-1
	for L < R {
		if s[L] == s[R] {
			L++
			R--
		} else {
			return false
		}
	}
	return true
}

func main() {

}
