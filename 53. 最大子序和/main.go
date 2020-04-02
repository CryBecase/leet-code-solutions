package main

// https://leetcode-cn.com/problems/maximum-subarray/

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := nums[0]
	sum := 0
	for _, num := range nums {
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}
		result = maxFunc(result, sum)
	}
	return result
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
