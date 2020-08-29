package main

// https://leetcode-cn.com/problems/shortest-palindrome/

func shortestPalindrome(s string) string {
	res := s
	for i := len(s) - 1; i >= 0; i-- {
		L, R := i, i
		str1 := ""
		str2 := ""

		if L-1 >= 0 && s[L-1] == s[R] {
			L -= 1
			for L >= 0 && R < len(s) && s[L] == s[R] {
				L--
				R++
			}
			if L == -1 {
				for j := len(s) - 1; j >= R; j-- {
					str1 += string(s[j])
				}
				if str1 == "" {
					return res
				}
			}
		}

		L, R = i, i
		for L >= 0 && R < len(s) && s[L] == s[R] {
			L--
			R++
		}
		if L == -1 {
			for j := len(s) - 1; j >= R; j-- {
				str2 += string(s[j])
			}
			if str2 == "" {
				return res
			}
		}

		if str1 != "" && str2 != "" {
			if len(str1) < len(str2) {
				return str1 + res
			} else {
				return str2 + res
			}
		}
		if str1 != "" {
			return str1 + res
		}
		if str2 != "" {
			return str2 + res
		}
	}
	return res
}

func main() {

}
