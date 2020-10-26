package main

// https://leetcode-cn.com/problems/how-many-numbers-are-smaller-than-the-current-number/

func smallerNumbersThanCurrent(nums []int) []int {
	counter := make([]int, 101)
	for _, n := range nums {
		counter[n]++
	}
	// 前缀和
	for i := range counter {
		if i == 0 {
			continue
		}
		counter[i] += counter[i-1]
	}
	res := make([]int, len(nums))
	for i, n := range nums {
		if n == 0 {
			res[i] = 0
		} else {
			res[i] = counter[n-1]
		}
	}
	return res
}

func main() {

}
