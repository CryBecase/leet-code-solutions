package main

import "sort"

// https://leetcode-cn.com/problems/assign-cookies/description/

// g是孩子 s是饼干
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	si, gi := 0, 0
	for si < len(s) && gi < len(g) {
		if g[gi] <= s[si] {
			gi++
		}
		si++
	}
	return gi
}

func main() {

}
