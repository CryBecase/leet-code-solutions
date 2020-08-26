package main

// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

var numMap = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	res := make([]string, 0)
	path := ""

	dfs := func(digits string, idx int) {}
	dfs = func(digits string, idx int) {
		if idx >= len(digits) {
			return
		}

		if len(path)+1 == len(digits) {
			for _, c := range numMap[digits[idx]] {
				temp := path
				temp += c
				res = append(res, temp)
			}
			return
		}

		temp := path
		for _, c := range numMap[digits[idx]] {
			path += c
			idx++
			dfs(digits, idx)
			idx--
			path = temp
		}
	}
	dfs(digits, 0)
	return res
}

func main() {

}
