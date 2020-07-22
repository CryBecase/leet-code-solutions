package main

// https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/

func minArray(numbers []int) int {
	L, R := 0, len(numbers)-1
	for L < R {
		mid := L + (R-L)/2
		if numbers[mid] > numbers[R] {
			L = mid + 1
		} else if numbers[mid] < numbers[R] {
			R = mid
		} else {
			R--
		}
	}
	return numbers[L]
}

func main() {

}
