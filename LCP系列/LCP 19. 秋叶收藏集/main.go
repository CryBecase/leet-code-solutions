package main

import "fmt"

// https://leetcode-cn.com/problems/UlBDOe/

func minimumOperations(leaves string) int {
	// dp[i][0] 表示到第 i 个字符是 r 格式，需要花费的最少次数
	// dp[i][1] 表示到第 i 个字符是 ry 格式，需要花费的最少次数
	// dp[i][2] 表示到第 i 个字符是 ryr 格式，需要花费的最少次数
	dp := make([][]int, len(leaves))
	for i := range dp {
		dp[i] = make([]int, 3)
	}
	dp[0][0] = r(leaves[0])
	for i := 1; i < len(leaves); i++ {
		leave := leaves[i]
		dp[i][0] = dp[i-1][0] + r(leave)
		dp[i][1] = dp[i-1][0] + y(leave)
		if i > 1 {
			dp[i][1] = minFunc(dp[i-1][0]+y(leave), dp[i-1][1]+y(leave))
			dp[i][2] = dp[i-1][1] + r(leave)
		}
		if i > 2 {
			dp[i][2] = minFunc(dp[i][2], dp[i-1][2]+r(leave))
		}
	}

	return dp[len(leaves)-1][2]
}

func r(c uint8) int {
	if c == 'r' {
		return 0
	}
	return 1
}

func y(c uint8) int {
	if c == 'y' {
		return 0
	}
	return 1
}

func minFunc(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	fmt.Println(minimumOperations("yry"))
}
