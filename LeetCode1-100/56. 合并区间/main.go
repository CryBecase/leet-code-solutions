package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/merge-intervals/

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	quickSort(intervals)

	var result [][]int

	for i := 0; i < len(intervals); i++ {
		L, R := intervals[i][0], intervals[i][1]
		for i+1 < len(intervals) && R >= intervals[i+1][0] {
			R = maxFunc(R, intervals[i+1][1])
			i++
		}
		result = append(result, []int{L, R})
	}

	return result
}

func quickSort(arr [][]int) {
	qs(arr, 0, len(arr)-1)
}

func qs(arr [][]int, L, R int) {
	if L < R {
		mid := L + (R-L)>>2
		arr[mid], arr[R] = arr[R], arr[mid]
		p := partition(arr, L, R)
		qs(arr, L, p[0])
		qs(arr, p[1], R)
	}
}

func partition(arr [][]int, L, R int) [2]int {
	less, more := L-1, R
	for L < more {
		if arr[L][0] < arr[R][0] {
			less++
			arr[less], arr[L] = arr[L], arr[less]
			L++
		} else if arr[L][0] > arr[R][0] {
			more--
			arr[more], arr[L] = arr[L], arr[more]
		} else {
			L++
		}
	}
	arr[more], arr[R] = arr[R], arr[more]
	more++
	return [2]int{less, more}
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	i := merge([][]int{
		{2, 3},
		{4, 5},
		{6, 7},
		{8, 9},
		{1, 10},
	})
	fmt.Println(i)
}
