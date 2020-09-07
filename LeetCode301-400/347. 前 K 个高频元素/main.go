package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/top-k-frequent-elements/description/

func topKFrequent(nums []int, k int) []int {
	frepMap := make(map[int]int)
	for _, num := range nums {
		frepMap[num]++
	}

	maxFrep := 0
	for _, frep := range frepMap {
		maxFrep = maxFunc(maxFrep, frep)
	}

	// arr[frep][num]
	bucket := make([][]int, maxFrep+1, maxFrep+1)
	for num, frep := range frepMap {
		bucket[frep] = append(bucket[frep], num)
	}

	result := make([]int, k, k)

	resultIndex := 0

	for i := len(bucket) - 1; i >= 0 && resultIndex != k; i-- {
		for j := 0; j < len(bucket[i]) && resultIndex != k; j++ {
			result[resultIndex] = bucket[i][j]
			resultIndex++
		}
	}

	return result
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	ints := topKFrequent([]int{1, 2}, 2)
	fmt.Println(ints)
}
