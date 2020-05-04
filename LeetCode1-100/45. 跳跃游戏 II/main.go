package main

import "fmt"

// https://leetcode-cn.com/problems/jump-game-ii/

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	L, count := 0, 0
	for {
		jumpLen := nums[L]
		if jumpLen == 0 {
			return 0
		}
		max := 0
		jumpTo := 0
		for i := 0; i < jumpLen && (L+i+1) < len(nums); i++ {
			if L+i+1+nums[L+i+1] >= max {
				max = L + i + 1 + nums[L+i+1]
				jumpTo = i + 1
			}
			if L+i+1 == len(nums)-1 {
				return count + 1
			}
		}
		L += jumpTo
		count++
		if L >= len(nums)-1 {
			return count
		}
	}
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(jump([]int{0, 2, 1}))
}
