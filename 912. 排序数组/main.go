package main

// https://leetcode-cn.com/problems/sort-an-array/

// 随机快排
func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	RandomQuickSort(nums, 0, len(nums)-1)
	return nums
}

func RandomQuickSort(nums []int, L, R int) {
	if L < R {
		r := L + (R-L)/2
		nums[r], nums[R] = nums[R], nums[r]
		p := partition(nums, L, R)
		RandomQuickSort(nums, L, p[0])
		RandomQuickSort(nums, p[1], R)
	}
}

func partition(nums []int, L, R int) [2]int {
	less, more := L-1, R
	for L < more {
		if nums[L] < nums[R] {
			less++
			nums[less], nums[L] = nums[L], nums[less]
			L++
		} else if nums[L] > nums[R] {
			more--
			nums[more], nums[L] = nums[L], nums[more]
		} else {
			L++
		}
	}
	nums[R], nums[more] = nums[more], nums[R]
	return [2]int{less, more + 1}
}

func main() {

}
