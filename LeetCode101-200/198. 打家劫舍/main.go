package main

// https://leetcode-cn.com/problems/house-robber/description/

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return maxFunc(nums[0], nums[1])
	}
	// dp[i] 表示第i个房间能够偷的最大值
	dp := make([]int, 0)
	dp = append(dp, 0)
	dp = append(dp, nums[0])
	dp = append(dp, maxFunc(nums[0], nums[1]))
	max := 0
	for i := 2; i < len(nums); i++ {
		// 这里的 i 是 index，所以本来的 dp[i-1] 现在是 dp[i]
		num := maxFunc(dp[i], dp[i-1]+nums[i])
		dp = append(dp, num)
		max = maxFunc(max, num)
	}
	return max
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
