package main

// https://leetcode-cn.com/problems/jewels-and-stones/

func numJewelsInStones(J string, S string) int {
	helper := make(map[rune]bool)
	for _, c := range J {
		helper[c] = true
	}

	res := 0
	for _, c := range S {
		if _, ok := helper[c]; ok {
			res++
		}
	}
	return res
}

func main() {

}
