package main

// 堆排序（最大堆）
// parent = (child-1) / 2
// left = parent * 2 + 1
// right = parent * 2 + 2
func sortArray3(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	for i := range nums {
		heapInsert(nums, i)
	}
	heapSize := len(nums)
	for heapSize > 0 {
		heapSize--
		nums[0], nums[heapSize] = nums[heapSize], nums[0]
		heapify(nums, heapSize)
	}
	return nums
}

// 往下走
func heapify(nums []int, size int) {
	parent := 0
	left := 1
	largest := 1
	for left < size {
		if left+1 < size && nums[left+1] > nums[left] {
			largest = left + 1
		}
		if nums[parent] >= nums[largest] {
			break
		} else {
			nums[parent], nums[largest] = nums[largest], nums[parent]
			parent = largest
			left = parent*2 + 1
			largest = left
		}
	}
}

// 往上走
func heapInsert(nums []int, i int) {
	for nums[i] > nums[(i-1)/2] {
		nums[i], nums[(i-1)/2] = nums[(i-1)/2], nums[i]
		i = (i - 1) / 2
	}
}

func main() {

}
