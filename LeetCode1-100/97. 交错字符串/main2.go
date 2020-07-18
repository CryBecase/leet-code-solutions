package main

// https://leetcode-cn.com/problems/interleaving-string/

func isInterleave2(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	// dp[i][j] 表示 s1 的前 i 个字符和 s2 的前 j 个字符能不能构成 s3 的前 i+j 个字符
	// dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
	dp := make([][]bool, len(s1)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}
	dp[0][0] = true
	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i > 0 {
				dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i+j-1]
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
			}
		}
	}
	return dp[len(s1)][len(s2)]
}

func main() {

}
