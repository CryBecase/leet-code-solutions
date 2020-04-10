package main

// https://leetcode-cn.com/problems/longest-common-prefix/

func longestCommonPrefix(strs []string) string {
	var result []byte
	if len(strs) == 0 {
		return string(result)
	}
	if len(strs) == 1 {
		return strs[0]
	}

	curIndex := 0
	for {
		if curIndex >= len(strs[0]) {
			return string(result)
		}
		curByte := strs[0][curIndex]
		for i := 1; i < len(strs); i++ {
			if curIndex >= len(strs[i]) {
				return string(result)
			}
			if curByte != strs[i][curIndex] {
				return string(result)
			}
		}
		result = append(result, curByte)
		curIndex++
	}
}

func main() {

}
