package main

// https://leetcode-cn.com/problems/the-masseuse-lcci/

func massage(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	var dp = make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = maxFunc(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		//这一次选择的话 dp[i-2]+num[i]
		//这一次不选择的话 dp[i-1]
		dp[i] = maxFunc(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
