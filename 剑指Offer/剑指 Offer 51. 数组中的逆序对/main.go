package main

import "fmt"

// https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/

func reversePairs(nums []int) int {
	res := mergeSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
	return res
}

func mergeSort(arr []int, L, R int) int {
	if L >= R {
		return 0
	}
	res := 0
	mid := L + (R-L)/2
	res += mergeSort(arr, L, mid)
	res += mergeSort(arr, mid+1, R)
	res += merge(arr, L, mid, R)
	return res
}

func merge(arr []int, L, mid, R int) int {
	res := 0

	LIdx, RIdx := L, mid+1
	tmp := make([]int, 0)
	for LIdx <= mid && RIdx <= R {
		if arr[LIdx] <= arr[RIdx] {
			tmp = append(tmp, arr[LIdx])
			LIdx++
		} else {
			res += (mid - LIdx) + 1
			tmp = append(tmp, arr[RIdx])
			RIdx++
		}
	}
	for LIdx <= mid {
		tmp = append(tmp, arr[LIdx])
		LIdx++
	}
	for RIdx <= R {
		tmp = append(tmp, arr[RIdx])
		RIdx++
	}
	for i := L; i <= R; i++ {
		arr[i] = tmp[i-L]
	}
	return res
}

func main() {
	reversePairs([]int{2, 54, 5, 12, 215, 6, 23, 567, 90})
}
