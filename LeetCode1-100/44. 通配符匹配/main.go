package main

// https://leetcode-cn.com/problems/wildcard-matching/

func isMatch(s string, p string) bool {
	// 将最后一个字符变成 * 号
	for len(s) > 0 && len(p) > 0 && p[len(p)-1] != '*' {
		if charMatch(s[len(s)-1], p[len(p)-1]) {
			s = s[:len(s)-1]
			p = p[:len(p)-1]
		} else {
			return false
		}
	}
	if len(p) == 0 {
		// p 中没有 * 号，且 s 和 p 的长度不一样
		return len(s) == 0
	}
	sIndex, pIndex := 0, 0
	sRecord, pRecord := -1, -1
	for sIndex < len(s) && pIndex < len(p) {
		if p[pIndex] == '*' {
			pIndex++
			sRecord, pRecord = sIndex, pIndex
		} else if charMatch(s[sIndex], p[pIndex]) {
			sIndex++
			pIndex++
		} else if sRecord != -1 && sRecord+1 < len(s) {
			sRecord++
			sIndex, pIndex = sRecord, pRecord
		} else {
			return false
		}
	}
	return allStars(p, pIndex, len(p))
}

func allStars(str string, left, right int) bool {
	for i := left; i < right; i++ {
		if str[i] != '*' {
			return false
		}
	}
	return true
}

func charMatch(u, v byte) bool {
	return u == v || v == '?'
}

func main() {

}
