package main

// https://leetcode-cn.com/problems/is-subsequence/

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}
	p1, p2 := 0, 0
	for ; p1 < len(t); p1++ {
		if t[p1] == s[p2] {
			p2++
			if p2 == len(s) {
				return true
			}
		}
	}
	return false
}

func main() {

}
