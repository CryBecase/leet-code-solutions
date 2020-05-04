package main

import "math"

// https://leetcode-cn.com/problems/sub-sort-lcci/

func subSort(array []int) []int {
	if array == nil || len(array) == 0 {
		return []int{-1, -1}
	}

	length := len(array)
	L, R := -1, -1
	min, max := math.MaxInt32, math.MinInt32
	for i, value := range array {
		if value < max {
			R = i
		} else {
			max = maxFunc(max, value)
		}
		if array[length-1-i] > min {
			L = length - 1 - i
		} else {
			min = minFunc(min, array[length-1-i])
		}
	}

	return []int{L, R}
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func minFunc(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {

}
