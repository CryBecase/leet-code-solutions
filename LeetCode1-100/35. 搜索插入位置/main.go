package main

import "fmt"

// https://leetcode-cn.com/problems/search-insert-position/

func searchInsert(nums []int, target int) int {
	length := len(nums)
	if target > nums[length-1] {
		return length
	}
	L, R := 0, length-1
	for L < R {
		mid := (R-L)/2 + L
		if target > nums[mid] {
			L = mid + 1
		} else {
			R = mid
		}
	}
	return L
}

func main() {
	fmt.Println(searchInsert([]int{1, 3}, 2))
}
