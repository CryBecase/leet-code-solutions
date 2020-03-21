package main

// https://leetcode-cn.com/problems/water-and-jug-problem/

// 见官方解答 方法二
func canMeasureWater(x int, y int, z int) bool {
	if x+y < z {
		return false
	}
	d := gcd(x, y)
	if d == 0 {
		return z == 0
	} else {
		return z%d == 0
	}
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func main() {
}
