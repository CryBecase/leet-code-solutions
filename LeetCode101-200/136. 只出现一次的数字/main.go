package main

// https://leetcode-cn.com/problems/single-number/

func singleNumber(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	helper := make(map[int]int)
	for _, num := range nums {
		if _, ok := helper[num]; ok {
			helper[num]++
		} else {
			helper[num] = 1
		}
	}

	for k, v := range helper {
		if v == 1 {
			return k
		}
	}

	return 0
}

func main() {

}
