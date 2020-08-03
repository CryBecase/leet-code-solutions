package main

import "fmt"

//面试题 08.03. 魔术索引
//https://leetcode-cn.com/problems/magic-index-lcci/

func findMagicIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if nums[0] > len(nums)-1 {
		return -1
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == i {
			return i
		}
		if nums[i] > i {
			i = nums[i] - 1
		}
	}
	// for i, num := range nums {
	//     if num == i {
	//         return i
	//     }
	// }
	return -1
}

func main() {
	nums := []int{0, 2, 3, 4, 5}
	result := findMagicIndex(nums)
	fmt.Println(result)
}
