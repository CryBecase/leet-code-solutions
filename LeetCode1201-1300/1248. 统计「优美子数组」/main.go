package main

import "fmt"

// https://leetcode-cn.com/problems/count-number-of-nice-subarrays/

func numberOfSubarrays(nums []int, k int) int {
	if len(nums) == k {
		return 1
	}
	if len(nums) < k {
		return 0
	}
	count := make([]int, len(nums)+1)
	count[0] = 1

	odd, result := 0, 0

	for i := 0; i < len(nums); i++ {
		odd += nums[i] & 1
		if odd >= k {
			result += count[odd-k]
		}
		count[odd] += 1
	}

	return result
}

func main() {
	fmt.Println(numberOfSubarrays([]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2))
}
