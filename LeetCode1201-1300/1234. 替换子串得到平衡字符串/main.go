package main

import "fmt"

// https://leetcode-cn.com/problems/replace-the-substring-for-balanced-string/

var m = map[uint8]int{
	'Q': 0,
	'W': 1,
	'E': 2,
	'R': 3,
}

func balancedString(s string) int {
	res := len(s)
	avg := len(s) / 4
	counter := make([]int, 4)
	for i := 0; i < len(s); i++ {
		counter[m[s[i]]]++
	}
	L, R := 0, 0
	for R < len(s) {
		counter[m[s[R]]]--
		for L < len(s) && counter[m['Q']] <= avg && counter[m['W']] <= avg && counter[m['E']] <= avg && counter[m['R']] <= avg {
			res = minFunc(res, R-L+1)
			counter[m[s[L]]]++
			L++
		}
		R++
	}
	return res
}

func minFunc(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	fmt.Println(balancedString("WWQQRRRRQRQQ"))
}
