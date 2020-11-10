package main

import "sort"

// https://leetcode-cn.com/problems/sort-integers-by-the-number-of-1-bits/

func sortByBits(arr []int) []int {
	type Nums struct {
		Num  int
		Ones int
	}
	nums := make([]Nums, 0)
	for _, n := range arr {
		nums = append(nums, Nums{n, oneCount(n)})
	}

	sort.Slice(nums, func(i, j int) bool {
		if nums[i].Ones < nums[j].Ones {
			return true
		} else if nums[i].Ones == nums[j].Ones {
			if nums[i].Num < nums[j].Num {
				return true
			}
			return false
		} else {
			return false
		}
	})

	res := make([]int, 0)
	for _, n := range nums {
		res = append(res, n.Num)
	}
	return res
}

// 二进制中 1 的个数
func oneCount(n int) int {
	if n == 0 {
		return 0
	}
	return oneCount(n/2) + n%2
}

func main() {

}
