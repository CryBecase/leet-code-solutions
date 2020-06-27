package main

// https://leetcode-cn.com/problems/first-missing-positive/

func firstMissingPositive(nums []int) int {
	for i := range nums {
		for nums[i] > 0 && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i, num := range nums {
		if num != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func main() {
	firstMissingPositive([]int{3, 4, -1, 1})
}
