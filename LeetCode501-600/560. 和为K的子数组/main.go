package main

// https://leetcode-cn.com/problems/subarray-sum-equals-k/

func subarraySum(nums []int, k int) int {
	L, R := 0, 0
	ans := 0
	for R < len(nums) {
		L = R
		sum := 0
		for L >= 0 {
			sum += nums[L]
			if sum == k {
				ans++
			}
			L--
		}
		R++
	}
	return ans
}

func main() {

}
