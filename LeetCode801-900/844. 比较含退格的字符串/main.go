package main

// https://leetcode-cn.com/problems/backspace-string-compare/

func backspaceCompare(S string, T string) bool {
	s := make([]rune, 0)
	for _, c := range S {
		if c == '#' {
			if len(s) != 0 {
				s = s[:len(s)-1]
			}
		} else {
			s = append(s, c)
		}
	}
	t := make([]rune, 0)
	for _, c := range T {
		if c == '#' {
			if len(t) != 0 {
				t = t[:len(t)-1]
			}
		} else {
			t = append(t, c)
		}
	}
	return string(s) == string(t)
}

func main() {

}
