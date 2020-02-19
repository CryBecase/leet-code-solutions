package main

/**
https://leetcode-cn.com/problems/container-with-most-water/
*/

// 双指针
// 小的指针往内移动
func maxArea(height []int) int {
	// result = right-left * max(left, right)
	result := 0
	left, right := 0, len(height)-1
	for left < right {
		if result < (right-left)*minFunc(height[left], height[right]) {
			result = (right - left) * minFunc(height[left], height[right])
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
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
