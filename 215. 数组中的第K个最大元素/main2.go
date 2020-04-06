package main

import (
	"fmt"
)

// 快速选择
func findKthLargest2(nums []int, k int) int {
	if len(nums) < k {
		return -1
	}
	// 找第 k 大的 = 找第 len-k 小的
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, L int, R int, k int) int {
	if L == R {
		return nums[L]
	}

	pivot := L + (R-L)/2
	nums[pivot], nums[R] = nums[R], nums[pivot]
	pivot = partition(nums, L, R)

	if pivot == k {
		return nums[pivot]
	} else if pivot < k {
		// right
		return quickSelect(nums, pivot+1, R, k)
	} else {
		// left
		return quickSelect(nums, L, pivot-1, k)
	}
}

func partition(nums []int, L, R int) int {
	less := L
	pivot := nums[R]
	for i := L; i <= R; i++ {
		if nums[i] < pivot {
			nums[less], nums[i] = nums[i], nums[less]
			less++
		}
	}
	nums[less], nums[R] = nums[R], nums[less]
	return less
}

func main() {
	test := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest2(test, 4))
}
