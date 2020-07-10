package main

/**
https://leetcode-cn.com/problems/container-with-most-water/
*/

// 双指针
// 小的指针往内移动
func maxArea(height []int) int {
	// result = R-L * min(L, R)
	result := 0
	L, R := 0, len(height)-1
	for L < R {
		if result < (R-L)*minFunc(height[L], height[R]) {
			result = (R - L) * minFunc(height[L], height[R])
		}
		if height[L] < height[R] {
			L++
		} else {
			R--
		}
	}
	return result
}

func minFunc(i, j int) int {
	if i <= j {
		return i
	} else {
		return j
	}
}

func main() {

}
