package main

// https://leetcode-cn.com/problems/trapping-rain-water/

func trap(height []int) int {
	res, leftMax, rightMax := 0, 0, 0
	for left, right := 0, len(height)-1; left < right; {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				// 右边必定有柱子挡水，所以，遇到所有值小于等于leftMax的，全部加入水池
				res += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				// 左边必定有柱子挡水，所以，遇到所有值小于等于rightMax的，全部加入水池
				res += rightMax - height[right]
			}
			right--
		}
	}
	return res
}

func main() {
	trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
}
