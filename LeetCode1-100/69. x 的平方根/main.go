package main

// https://leetcode-cn.com/problems/sqrtx/

func mySqrt(x int) int {
	L, R := 0, x
	res := -1
	for L <= R {
		mid := L + (R-L)/2
		if mid*mid <= x {
			res = mid
			L = mid + 1
		} else {
			R = mid - 1
		}
	}
	return res
}

func main() {

}
