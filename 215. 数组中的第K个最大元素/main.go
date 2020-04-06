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
	if len(nums) < k {
		return -1
	}
	minHeap := make([]int, 0, k)

	for i := 0; i < k; i++ {
		minHeap = append(minHeap, nums[i])
	}

	for i := 0; i < k; i++ {
		heapInsert(minHeap, i)
	}

	for i := k; i < len(nums); i++ {
		if nums[i] > minHeap[0] {
			minHeap[0] = nums[i]
			heapify(minHeap)
		}
	}

	return minHeap[0]
}

func heapInsert(minHeap []int, i int) {
	for minHeap[i] < minHeap[(i-1)/2] {
		minHeap[i], minHeap[(i-1)/2] = minHeap[(i-1)/2], minHeap[i]
		i = (i - 1) / 2
	}
}

func heapify(minHeap []int) {
	root := 0
	left := 1
	least := 1
	for left < len(minHeap) {
		if left+1 < len(minHeap) && minHeap[left+1] < minHeap[left] {
			least = left + 1
		}
		if minHeap[least] >= minHeap[root] {
			break
		} else {
			minHeap[least], minHeap[root] = minHeap[root], minHeap[least]
			root = least
			left = root*2 + 1
			least = left
		}
	}
}

func main() {
	test := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(test, 4))
}
