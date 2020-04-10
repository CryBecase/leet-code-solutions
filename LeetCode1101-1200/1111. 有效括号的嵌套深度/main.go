package main

// https://leetcode-cn.com/problems/maximum-nesting-depth-of-two-valid-parentheses-strings/

func maxDepthAfterSplit(seq string) []int {
	helper := 0
	result := make([]int, 0, len(seq))
	for _, c := range seq {
		if c == '(' {
			helper++
			result = append(result, helper%2)
		} else {
			result = append(result, helper%2)
			helper--
		}
	}
	return result
}

func main() {

}
