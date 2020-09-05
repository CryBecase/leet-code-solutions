package main

// https://leetcode-cn.com/problems/predict-the-winner/

// 区间 DP
func PredictTheWinner(nums []int) bool {
	// dp[i][j] 表示先手此区间 nums[i...j] 先手玩家对后手玩家的最大净胜分
	// 则 dp[0][len(nums)-1] >=0 就是结果
	dp := make([][]int, len(nums))

	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums))
		dp[i][i] = nums[i]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			// max(拿左边的时候剩下区间对手的最大净胜分, 拿右边的时候剩下区间对手的最大净胜分)
			dp[i][j] = maxFunc(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}

	return dp[0][len(nums)-1] >= 0
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
