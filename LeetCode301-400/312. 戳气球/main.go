package main

// https://leetcode-cn.com/problems/burst-balloons/

func maxCoins(nums []int) int {
	// dp[i][j] 表示从第i+1个气球到第j+1个气球能获取的最大硬币数
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, len(nums))
	}
	return helper(nums, dp, 0, len(nums)-1)
}

func helper(nums []int, dp [][]int, i, j int) int {
	if i > j {
		return 0
	}
	if dp[i][j] > 0 {
		return dp[i][j]
	}

	for k := i; k <= j; k++ {
		// 第k+1个气球在最后点爆
		L := helper(nums, dp, i, k-1)
		R := helper(nums, dp, k+1, j)
		coins := nums[k]
		if i-1 >= 0 {
			coins *= nums[i-1]
		}
		if j+1 < len(nums) {
			coins *= nums[j+1]
		}
		dp[i][j] = maxFunc(dp[i][j], L+coins+R)
	}
	return dp[i][j]
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
