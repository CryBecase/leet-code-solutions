package main

// https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self/

type Num struct {
	Value int
	Index int
}

func countSmaller(nums []int) []int {
	res := make([]int, len(nums))
	ns := make([]Num, 0)
	for i, num := range nums {
		ns = append(ns, Num{
			Value: num,
			Index: i,
		})
	}
	mergeSort(ns, 0, len(ns)-1, res)
	return res
}

func mergeSort(arr []Num, L, R int, res []int) {
	if L >= R {
		return
	}
	mid := L + (R-L)/2
	mergeSort(arr, L, mid, res)
	mergeSort(arr, mid+1, R, res)
	merge(arr, L, mid, R, res)
}

func merge(arr []Num, L, mid, R int, res []int) {
	LIdx, RIdx := L, mid+1
	tmp := make([]Num, 0)
	for LIdx <= mid && RIdx <= R {
		if arr[LIdx].Value > arr[RIdx].Value {
			tmp = append(tmp, arr[LIdx])
			res[arr[LIdx].Index] += R - RIdx + 1
			LIdx++
		} else {
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
}

func main() {
	countSmaller([]int{5, 2, 6, 1})
}
