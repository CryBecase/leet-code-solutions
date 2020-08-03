package main

import "math"

// https://leetcode-cn.com/problems/smallest-range-covering-elements-from-k-lists/

func smallestRange(nums [][]int) []int {
	size := len(nums)
	indices := map[int][]int{}
	xMin, xMax := math.MaxInt32, math.MinInt32 // 二分查找的区间
	for i := 0; i < size; i++ {
		for _, x := range nums[i] {
			indices[x] = append(indices[x], i)
			xMin = minFunc(xMin, x)
			xMax = maxFunc(xMax, x)
		}
	}
	window := make([]int, size)
	valid := 0
	L, R := xMin, xMin-1
	start, end := xMin, xMax
	for R < xMax {
		R++
		if len(indices[R]) > 0 {
			for _, i := range indices[R] {
				window[i]++
				if window[i] == 1 {
					valid++
				}
			}
			for valid == size {
				if R-L < end-start {
					start, end = L, R
				}
				for _, x := range indices[L] {
					window[x]--
					if window[x] == 0 {
						valid--
					}
				}
				L++
			}
		}
	}
	return []int{start, end}
}

func minFunc(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
