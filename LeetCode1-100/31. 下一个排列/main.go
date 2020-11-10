package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/next-permutation/

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	small := -1
	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i] > nums[i-1] {
			small = i - 1
			break
		}
	}

	if small != -1 {
		for i := len(nums) - 1; i >= small+1; i-- {
			if nums[i] > nums[small] {
				nums[i], nums[small] = nums[small], nums[i]
				break
			}
		}
	}

	L, R := small+1, len(nums)-1
	for L < R {
		nums[L], nums[R] = nums[R], nums[L]
		L++
		R--
	}
}

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums)
}
