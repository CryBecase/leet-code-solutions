package main

import "math"

// https://leetcode-cn.com/problems/reverse-integer

func reverse(x int) int {
	result := 0
	for x != 0 {
		pop := x % 10
		x /= 10

		// max32 个位数是 7
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && pop > 7) {
			return 0
		}
		// min32 个位数是 8
		if result < math.MinInt32/10 || (result == math.MinInt32/10 && pop < -8) {
			return 0
		}

		result = result*10 + pop
	}
	return result
}

func main() {

}
