package main

// https://leetcode-cn.com/problems/find-words-that-can-be-formed-by-characters/

func countCharacters(words []string, chars string) int {
	var byteCount [26]int
	for _, char := range chars {
		byteCount[char-'a']++
	}

	ret := 0
	for _, word := range words {
		bc, match := byteCount, true
		for _, char := range word {
			if bc[char-'a'] <= 0 {
				match = false
				break
			}
			bc[char-'a']--
		}
		if match {
			ret += len(word)
		}
	}
	return ret
}
func main() {
}
