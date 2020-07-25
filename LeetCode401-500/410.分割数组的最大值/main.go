package main

// https://leetcode-cn.com/problems/split-array-largest-sum/

func splitArray(nums []int, m int) int {
	L, R := 0, 0
	for _, num := range nums {
		L = maxFunc(L, num)
		R += num
	}
	for L < R {
		mid := L + (R-L)/2
		sp := split(nums, mid)
		if sp > m {
			L = mid + 1
		} else {
			R = mid
		}
	}
	return L
}

func split(nums []int, max int) int {
	sum := 0
	result := 1
	for _, num := range nums {
		sum += num
		if sum > max {
			sum = num
			result++
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

}
