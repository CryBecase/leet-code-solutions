package main

// https://leetcode-cn.com/problems/single-number-iii/

func singleNumber(nums []int) []int {
	helper := make(map[int]int)
	for _, num := range nums {
		if _, ok := helper[num]; ok {
			helper[num]++
		} else {
			helper[num] = 1
		}
	}

	res := make([]int, 0)
	for k, v := range helper {
		if v == 1 {
			res = append(res, k)
		}
	}

	return res
}

func main() {

}
