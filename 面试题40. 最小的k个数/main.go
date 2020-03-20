package main

// https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/

// 利用最大堆 每次删除最大的数
// left = root * 2 + 1
// right = root * 2 + 2
// root = (left-1) / 2 = (right-1) / 2
func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	maxHeap := make([]int, k, k)
	// init heap
	for i := 0; i < k; i++ {
		maxHeap[i] = arr[i]
	}
	// heapInsert
	for i := 0; i < k; i++ {
		heapInsert(maxHeap, i)
	}
	for i := k; i < len(arr); i++ {
		if arr[i] < maxHeap[0] {
			maxHeap[0] = arr[i]
			heapify(maxHeap)
		}
	}
	return maxHeap
}

func heapInsert(heap []int, index int) {
	for heap[index] > heap[(index-1)/2] {
		heap[index], heap[(index-1)/2] = heap[(index-1)/2], heap[index]
		index = (index - 1) / 2
	}
}

func heapify(heap []int) {
	root := 0
	left := root*2 + 1
	for left < len(heap) {
		largest := left
		if len(heap) > left+1 && heap[left+1] > heap[left] {
			largest = left + 1
		}
		if heap[root] > heap[largest] {
			break
		} else {
			heap[root], heap[largest] = heap[largest], heap[root]
			root = largest
			left = root*2 + 1
		}
	}
}

func main() {

}
