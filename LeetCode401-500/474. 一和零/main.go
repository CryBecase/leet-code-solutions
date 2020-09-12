package main

// https://leetcode-cn.com/problems/ones-and-zeroes/

func findMaxForm(strs []string, m int, n int) int {
	// dp[i][j][k] 表示到第 i 个字符串，使用 j 个 0，k 个 1 的时候能存放的字符串的最大个数
	dp := make([][][]int, len(strs)+1)
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}
	for i := 1; i <= len(strs); i++ {
		str := strs[i-1]
		zeroOneNums := getZeroAndOneNums(str)
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				dp[i][j][k] = dp[i-1][j][k]
				zeros := zeroOneNums[0]
				ones := zeroOneNums[1]
				if j >= zeros && k >= ones {
					dp[i][j][k] = maxFunc(dp[i-1][j][k], 1+dp[i-1][j-zeros][k-ones])
				}
			}
		}
	}
	return dp[len(strs)][m][n]
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func getZeroAndOneNums(str string) [2]int {
	res := [2]int{}
	for _, c := range str {
		if c-'0' == 0 {
			res[0]++
		} else {
			res[1]++
		}
	}
	return res
}

func main() {

}
