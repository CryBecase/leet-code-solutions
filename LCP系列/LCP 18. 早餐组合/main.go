package main

import "sort"

// https://leetcode-cn.com/problems/2vYnGI/

func breakfastNumber(staple []int, drinks []int, x int) int {
	sort.Ints(staple)
	sort.Ints(drinks)
	res := 0
	pt := len(drinks) - 1
	for i := range staple {
		if staple[i] > x {
			break
		}
		for pt >= 0 && staple[i]+drinks[pt] > x {
			pt--
		}
		res += pt + 1
	}

	return res % 1000000007
}

func main() {

}
