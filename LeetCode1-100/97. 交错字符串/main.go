package main

// https://leetcode-cn.com/problems/interleaving-string/

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	if len(s1) == 0 && len(s2) == 0 {
		return len(s3) == 0
	}
	if len(s1) == 0 {
		if len(s2) != len(s3) {
			return false
		}
		for i := 0; i < len(s2); i++ {
			if s2[i] != s3[i] {
				return false
			}
		}
		return true
	}
	if len(s2) == 0 {
		if len(s1) != len(s3) {
			return false
		}
		for i := 0; i < len(s1); i++ {
			if s1[i] != s3[i] {
				return false
			}
		}
		return true
	}
	if s1[0] == s2[0] && s1[0] == s3[0] {
		return isInterleave(s1[1:], s2, s3[1:]) || isInterleave(s1, s2[1:], s3[1:])
	} else if s1[0] == s3[0] {
		return isInterleave(s1[1:], s2, s3[1:])
	} else if s2[0] == s3[0] {
		return isInterleave(s1, s2[1:], s3[1:])
	}
	return false
}

func main() {

}
