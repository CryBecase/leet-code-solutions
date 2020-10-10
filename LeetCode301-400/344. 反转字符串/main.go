package main

// https://leetcode-cn.com/problems/reverse-string/

func reverseString(s []byte) {
	L, R := 0, len(s)-1
	for L < R {
		s[L], s[R] = s[R], s[L]
		L++
		R--
	}
}

func main() {

}
