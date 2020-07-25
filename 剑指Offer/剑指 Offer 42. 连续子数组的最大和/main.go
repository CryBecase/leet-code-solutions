package main

// https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// dp[i] 表示已第i+1个数为结尾的子数组的最大值
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	max := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = maxFunc(0, dp[i-1]) + nums[i]
		max = maxFunc(max, dp[i])
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
