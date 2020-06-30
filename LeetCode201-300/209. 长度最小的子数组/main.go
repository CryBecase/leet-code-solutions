package main

// https://leetcode-cn.com/problems/minimum-size-subarray-sum/

func minSubArrayLen(s int, nums []int) int {
	L, R := 0, 0
	ans := len(nums) + 1
	sum := 0
	for R < len(nums) {
		if sum+nums[R] < s {
			sum += nums[R]
			R++
		} else {
			sum -= nums[L]
			if R-L+1 < ans {
				ans = R - L + 1
			}
			L++
		}
	}
	if ans == len(nums)+1 {
		ans = 0
	}
	return ans
}

func main() {

}
