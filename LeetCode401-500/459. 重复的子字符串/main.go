package main

// https://leetcode-cn.com/problems/repeated-substring-pattern/

func repeatedSubstringPattern(s string) bool {
	ss := s + s
	return kmp(ss[1:len(ss)-1], s)
}

func kmp(query, pattern string) bool {
	n, m := len(query), len(pattern)

	next := make([]int, m)
	for i := 0; i < m; i++ {
		next[i] = -1
	}
	for i := 1; i < m; i++ {
		j := next[i-1]
		for j != -1 && pattern[j+1] != pattern[i] {
			j = next[j]
		}
		if pattern[j+1] == pattern[i] {
			next[i] = j + 1
		}
	}
	match := -1
	for i := 0; i < n; i++ {
		for match != -1 && pattern[match+1] != query[i] {
			match = next[match]
		}
		if pattern[match+1] == query[i] {
			match++
			if match == m-1 {
				return true
			}
		}
	}
	return false
}

func main() {

}
