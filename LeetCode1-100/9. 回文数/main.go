package main

/**
https://leetcode-cn.com/problems/palindrome-number/

判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
*/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	} else if x%10 == 0 {
		return false
	} else {
		pop, y := 0, 0
		for y < x {
			pop = x % 10
			x = x / 10
			y = y*10 + pop
		}
		if y > x {
			y = y / 10
			if y == x {
				return true
			} else {
				return false
			}
		} else {
			return true
		}
	}
}

func main() {

}
