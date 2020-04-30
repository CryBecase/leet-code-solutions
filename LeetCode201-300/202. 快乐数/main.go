package main

import "fmt"

// https://leetcode-cn.com/problems/happy-number/

func isHappy(n int) bool {
	helper := make(map[int]struct{})
	for {
		if n == 1 {
			return true
		}
		// 重复
		if _, ok := helper[n]; ok {
			return false
		}
		helper[n] = struct{}{}
		n = getNext(n)
	}
}

func getNext(n int) int {
	res := 0
	for n != 0 {
		res += (n % 10) * (n % 10)
		n /= 10
	}
	return res
}

func main() {
	next := getNext(2)
	fmt.Println(next)
}
