package main

import "sort"

// https://leetcode-cn.com/problems/relative-sort-array/

func relativeSortArray(arr1 []int, arr2 []int) []int {
	counter := make([]int, len(arr2))
	dic := make(map[int]int)
	for i, n := range arr2 {
		dic[n] = i
	}
	arr3 := make([]int, 0)
	for _, n := range arr1 {
		if idx, ok := dic[n]; ok {
			counter[idx]++
		} else {
			arr3 = append(arr3, n)
		}
	}
	sort.Ints(arr3)
	arr1 = arr1[0:0]
	for idx, count := range counter {
		for i := 0; i < count; i++ {
			arr1 = append(arr1, arr2[idx])
		}
	}
	for _, n := range arr3 {
		arr1 = append(arr1, n)
	}
	return arr1
}

func main() {

}
