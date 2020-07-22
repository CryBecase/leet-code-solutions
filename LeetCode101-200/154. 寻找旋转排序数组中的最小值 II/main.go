package main

// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/

func findMin(nums []int) int {
	L, R := 0, len(nums)-1
	for L < R {
		mid := L + (R-L)/2
		if nums[mid] > nums[R] {
			L = mid + 1
		} else if nums[mid] < nums[R] {
			R = mid
		} else {
			R--
		}
	}
	return nums[L]
}

func main() {

}
