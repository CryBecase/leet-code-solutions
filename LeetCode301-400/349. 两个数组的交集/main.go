package main

// https://leetcode-cn.com/problems/intersection-of-two-arrays/

func intersection(nums1 []int, nums2 []int) []int {
	helper := make(map[int]int)
	for _, n := range nums1 {
		helper[n] = 1
	}

	res := make([]int, 0)
	for _, n := range nums2 {
		if helper[n] == 1 {
			res = append(res, n)
			helper[n]++
		}
	}

	return res
}

func main() {

}
