package main

// https://leetcode-cn.com/problems/sliding-window-maximum/

func maxSlidingWindow(nums []int, k int) []int {
	result := []int{}
	queue := []int{}
	for i, num := range nums {
		j := len(queue) - 1
		for ; j >= 0 && nums[queue[j]] <= num; j-- {
		}
		queue = queue[0 : j+1]
		queue = append(queue, i)

		if i-queue[0] >= k {
			queue = queue[1:]
		}

		if i >= k-1 {
			result = append(result, nums[queue[0]])
		}
	}
	return result
}

func main() {

}
