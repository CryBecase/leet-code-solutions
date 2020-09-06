package main

import (
	"fmt"
	"math"
)

func minCost(s string, cost []int) int {
	if len(s) < 2 {
		return 0
	}
	res := 0
	L, R := 0, 1
	isSame := false
	for R < len(s) {
		if s[L] == s[R] {
			R++
			isSame = true
		} else {
			if isSame {
				max := math.MinInt32
				sum := 0
				for i := L; i <= R-1; i++ {
					max = maxFunc(cost[i], max)
					sum += cost[i]
				}
				sum -= max
				res += sum
				isSame = false
			}
			L = R
			R++
		}
	}
	if isSame {
		max := math.MinInt32
		sum := 0
		for i := L; i <= R-1; i++ {
			max = maxFunc(cost[i], max)
			sum += cost[i]
		}
		sum -= max
		res += sum
		isSame = false
	}

	return res
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(minCost("aaabbbabbbb", []int{3, 5, 10, 7, 5, 3, 5, 5, 4, 8, 1}))
	// 8 + 8 + 10 = 26
}
