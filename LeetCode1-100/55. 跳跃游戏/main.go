package main

// https://leetcode-cn.com/problems/jump-game/

func canJump(nums []int) bool {
	maxIndex := 0
	for i, num := range nums {
		if i > maxIndex {
			// 跳不到
			return false
		}
		maxIndex = maxFunc(maxIndex, i+num)
	}
	return true
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
