package main

import "fmt"

// https://leetcode-cn.com/problems/super-egg-drop/

// https://www.bilibili.com/video/BV1KE41137PK

// dp[n][k] 表示 有 n 层楼，有 k 个鸡蛋，最差情况下最少扔多少次能确定
func superEggDrop(K int, N int) int {
	if N == 1 {
		return 1
	}
	dp := make([][]int, N+1)
	for n := range dp {
		dp[n] = make([]int, K+1)
	}

	for k := 1; k <= K; k++ {
		dp[1][k] = 1
	}

	for n := 1; n <= N; n++ {
		dp[n][1] = n
	}

	for n := 2; n <= N; n++ {
		for k := 2; k <= K; k++ {
			//dp[n][k] = 99999 // 赋初始值，因为接下来是求 min 的过程
			// 线性搜索会超时
			//for i := 1; i <= n; i++ {
			//	dp[n][k] = minFunc(dp[n][k], maxFunc(dp[i-1][k-1], dp[n-i][k])+1)
			//}
			// 二分搜索
			// 这时候求二者的较大值，再求这些最大值之中的最小值，其实就是求这两条直线交点，也就是红色折线的最低点嘛。
			// dp[i-1][k-1] 上升  dp[n-i][k] 下降
			left, right := 1, n
			for left < right {
				mid := left + (right-left+1)/2
				broke := dp[mid-1][k-1]  // 碎
				notBroke := dp[n-mid][k] // 不碎
				if broke > notBroke {
					right = mid - 1
				} else {
					left = mid
				}
			}
			dp[n][k] = maxFunc(dp[left-1][k-1], dp[n-left][k]) + 1
		}
	}

	return dp[N][K]
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(superEggDrop(3, 14))
}
