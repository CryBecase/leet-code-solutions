package main

import "fmt"

/**
https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

示例 1:

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4
说明:

你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
*/

// 左孩子 = 根 * 2
// 右孩子 = 根 * 2 + 1
// 根 = 孩子 / 2

func findKthLargest(nums []int, k int) int {
	if k > len(nums) {
		return -1
	}

	maxHeap := make([]int, len(nums))
	copy(maxHeap, nums)
	for i := 0; i < len(nums); i++ {
		heapInsert(maxHeap, i)
	}

	size := len(maxHeap)
	for i := 0; i < k; i++ {
		maxHeap[0], maxHeap[size-1] = maxHeap[size-1], maxHeap[0]
		heapify(maxHeap, size-1)
		size--
	}

	return maxHeap[len(maxHeap)-k]
}

func heapInsert(maxHeap []int, i int) {
	for maxHeap[i] > maxHeap[(i-1)/2] {
		maxHeap[i], maxHeap[(i-1)/2] = maxHeap[(i-1)/2], maxHeap[i]
		i = (i - 1) / 2
	}
}

func heapify(maxHeap []int, size int) {
	index := 0
	left := index*2 + 1
	for left < size {
		largest := left
		if left+1 < size && maxHeap[left+1] > maxHeap[left] {
			largest = left + 1
		}
		if maxHeap[largest] < maxHeap[index] {
			largest = index
			break
		} else {
			maxHeap[largest], maxHeap[index] = maxHeap[index], maxHeap[largest]
			index = largest
			left = index*2 + 1
		}
	}
}

func main() {
	test := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(test, 4))
}
