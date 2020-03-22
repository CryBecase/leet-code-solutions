package main

import "fmt"

// https://leetcode-cn.com/problems/minimum-increment-to-make-array-unique/

func minIncrementForUnique(A []int) int {
	max := 0
	for _, v := range A {
		max = maxFunc(v, max)
	}

	bucket := make([]int, max+1+len(A), max+1+len(A))
	for _, v := range A {
		bucket[v]++
	}

	result, taken := 0, 0
	for i := 0; i < len(bucket); i++ {
		if bucket[i] > 1 {
			taken += bucket[i] - 1
			result -= i * (bucket[i] - 1)
		} else if taken > 0 && bucket[i] == 0 {
			taken--
			result += i
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
	fmt.Println(minIncrementForUnique([]int{1, 1, 2, 0}))
}
