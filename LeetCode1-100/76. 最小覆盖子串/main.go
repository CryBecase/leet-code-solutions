package main

// https://leetcode-cn.com/problems/minimum-window-substring/

func minWindow(s string, t string) string {
	window, need := make(map[uint8]int), make(map[uint8]int)
	for i := range t {
		need[t[i]]++
	}
	L, R := 0, 0                 // 左右指针 左闭右开 [L,R)
	start, length := 0, len(s)+1 // 记录结果
	valid := 0                   // 记录窗口中有多少需要的结果
	for R < len(s) {
		R++
		char := s[R-1]
		if need[char] > 0 {
			// 需要
			window[char]++
			if window[char] == need[char] {
				valid++
			}
		}
		for valid == len(need) {
			if R-L < length {
				length = R - L
				start = L
			}
			d := s[L]
			L++
			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	if length == len(s)+1 {
		return ""
	}
	return s[start : start+length]
}

func main() {

}
