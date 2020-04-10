package main

import "strings"

// https://leetcode-cn.com/problems/integer-to-roman/

func intToRoman(num int) string {
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	b := strings.Builder{}
	for i := 0; i < 13; i++ {
		for num >= nums[i] {
			b.WriteString(romans[i])
			num -= nums[i]
		}
	}
	return b.String()
}

func main() {

}
