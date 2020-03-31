package main

// 归并排序
func sortArray2(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	MergeSort(nums, 0, len(nums)-1)
	return nums
}

func MergeSort(nums []int, L, R int) {
	if L < R {
		mid := L + (R-L)/2
		MergeSort(nums, L, mid)
		MergeSort(nums, mid+1, R)
		merge(nums, L, mid, R)
	}
}

func merge(nums []int, L, mid, R int) {
	helper := make([]int, R-L+1, R-L+1)
	i, j := L, mid+1
	index := 0
	for i <= mid && j <= R {
		if nums[i] <= nums[j] {
			helper[index] = nums[i]
			i++
		} else {
			helper[index] = nums[j]
			j++
		}
		index++
	}
	for i <= mid {
		helper[index] = nums[i]
		index++
		i++
	}
	for j <= R {
		helper[index] = nums[j]
		index++
		j++
	}
	for i, v := range helper {
		nums[L+i] = v
	}
}

func main() {

}
