package main

// https://leetcode-cn.com/problems/short-encoding-of-words/

func minimumLengthEncoding(words []string) int {
	result := 0
	helper := make(map[string]int)
	for _, word := range words {
		if _, ok := helper[word]; !ok {
			helper[word] = 1
		}
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
