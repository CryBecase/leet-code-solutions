package main

import "math"

/**
https://leetcode-cn.com/problems/sum-of-square-numbers
给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a^2 + b^2 = c。

示例1:

输入: 5
输出: True
解释: 1 * 1 + 2 * 2 = 5


示例2:

输入: 3
输出: False
*/

func judgeSquareSum(c int) bool {
	if c == 1 || c == 2 || c == 0 {
		return true
	}
	small, big := 0, int(math.Sqrt(float64(c)))
	for small <= big {
		if small*small+big*big < c {
			small++
		} else if small*small+big*big == c {
			return true
		} else {
			big--
		}
	}
	return false
}

func main() {

}
