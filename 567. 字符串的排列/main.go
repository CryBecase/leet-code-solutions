package main

// https://leetcode-cn.com/problems/permutation-in-string/

// map + 滑动窗口
func checkInclusion(s1 string, s2 string) bool {
	// init m1
	m1 := make(map[byte]interface{}, len(s1))
	for i := range s1 {
		m1[s1[i]] = nil
	}

	for i, j := 0, 0; j < len(s2); j++ {

	}
}

func main() {

}
