package main

import (
	"fmt"
)

/**
https://leetcode-cn.com/problems/top-k-frequent-elements/description/
给定一个非空的整数数组，返回其中出现频率前 k 高的元素。

示例 1:

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:

输入: nums = [1], k = 1
输出: [1]
说明：

你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。
*/

func topKFrequent(nums []int, k int) []int {
	// map统计频率 key:int value value:frequent
	fi := make(map[int]int)
	for i := range nums {
		fi[nums[i]]++
	}

	max := 0
	for i := range fi {
		max = maxFunc(max, fi[i])
	}
	// 一维数组的下标是频率 一维数组里的数组是频率相同的数
	arr := make([][]int, max+1, max+1)
	for key := range fi {
		frequent := fi[key]
		arr[frequent] = append(arr[frequent], key)
	}

	// 从二维数组中拿到数
	result := make([]int, k, k)
X:
	for i, resultIndex := len(arr)-1, 0; i >= 0; i-- {
		for j := 0; j < len(arr[i]); j++ {
			result[resultIndex] = arr[i][j]
			resultIndex++
			if resultIndex == k {
				break X
			}
		}
	}

	return result
}

func maxFunc(i int, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func main() {
	ints := topKFrequent([]int{1, 2}, 2)
	fmt.Println(ints)
}
