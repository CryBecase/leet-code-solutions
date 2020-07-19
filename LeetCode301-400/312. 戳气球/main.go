package main

// https://leetcode-cn.com/problems/burst-balloons/

func maxCoins(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// dp[i][j] 表示戳破从第i+1个气球到第j+1个气球能获取的最大硬币数
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, len(nums))
	}
	for i := len(nums) - 1; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			for k := i; k <= j; k++ {
				// 选择第k+1个气球留到最后戳破

				// 从第i+1个气球到第k个气球被戳破，从第k+2个气球到第j+1个气球被戳破时，戳破第k+1个气球时获得的硬币数
				coinK := nums[k]
				if i-1 >= 0 {
					coinK *= nums[i-1]
				}
				if j+1 < len(nums) {
					coinK *= nums[j+1]
				}

				// 当第k+1个气球留到最后被戳破的情况下，戳破第i+1到第j+1的所有气球获得的硬币数
				coins := coinK
				if k-1 >= 0 {
					coins += dp[i][k-1]
				}
				if k+1 < len(nums) {
					coins += dp[k+1][j]
				}
				dp[i][j] = maxFunc(dp[i][j], coins)
			}

		}
	}
	return dp[0][len(nums)-1]
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
