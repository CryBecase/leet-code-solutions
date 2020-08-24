package main

// https://leetcode-cn.com/problems/repeated-substring-pattern/

func repeatedSubstringPattern(s string) bool {
	if len(s) == 1 {
		return false
	}

	if len(s) == 2 {
		return s[0] == s[1]
	}

	L, R := 0, 0
	for R < len(s)-1 {
		if s[R] == s[len(s)-1] {
			if repeated(s[L:R+1], s) {
				return true
			}
		}

		R++
	}
	return false
}

func repeated(subStr, str string) bool {
	if len(subStr) > len(str) {
		return false
	}
	if len(subStr) == len(str) {
		return subStr == str
	}
	if len(str)%len(subStr) != 0 {
		// str 与 subStr 不是整倍关系 则一定不是
		return false
	}

	p1, p2 := 0, len(subStr)
	for p2 < len(str) {
		if str[p2] != subStr[p1] {
			return false
		}

		p1++
		p2++
		if p1 == len(subStr) {
			p1 = 0
		}
	}
	return true
}

func main() {

}
