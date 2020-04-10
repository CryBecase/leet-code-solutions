package main

// https://leetcode-cn.com/problems/short-encoding-of-words/

func minimumLengthEncoding(words []string) int {
	result := 0
	helper := make(map[string]bool, len(words))
	for _, word := range words {
		helper[word] = true
	}
	for word := range helper {
		for i := 1; i < len(word); i++ {
			target := word[i:]
			if _, ok := helper[target]; ok {
				delete(helper, target)
			}
		}
	}
	for word := range helper {
		result += len(word) + 1
	}

	return result
}

func main() {

}
