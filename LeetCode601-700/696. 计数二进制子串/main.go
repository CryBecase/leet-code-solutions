package main

// https://leetcode-cn.com/problems/count-binary-substrings/

func countBinarySubstrings(s string) int {
	if s == "" {
		return 0
	}

	result := 0
	counts := make([]int, 0)
	p := 0
	for p < len(s) {
		cur := s[p]
		count := 0
		for p < len(s) && s[p] == cur {
			count++
			p++
		}
		counts = append(counts, count)
	}
	for i := 1; i < len(counts); i++ {
		result += minFunc(counts[i], counts[i-1])
	}
	return result
}

func minFunc(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {

}
