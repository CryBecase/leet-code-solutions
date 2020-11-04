package main

// https://leetcode-cn.com/problems/sliding-window-maximum/

// 每当数字 n rpush 就 rpop 比 n 小的所有数
// 保证 queue 中最左边的数是当前 queue 中最大的数
func maxSlidingWindow(nums []int, k int) []int {
	result := []int{}
	queue := []int{}
	for i, num := range nums {
		j := len(queue) - 1
		for ; j >= 0 && nums[queue[j]] <= num; j-- {
			// rpop
		}
		queue = queue[0 : j+1]

		// rpush
		queue = append(queue, i)

		// queue 超过了窗口大小
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
