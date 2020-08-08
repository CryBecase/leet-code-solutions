package main

import "sort"

// https://leetcode-cn.com/problems/3sum/

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	result := make([][]int, 0)
	sort.Ints(nums)
	for i, num := range nums {
		if num > 0 {
			break
		}
		if i > 0 && num == nums[i-1] {
			// 与上一个相同，则跳过，避免相同答案
			continue
		}

		L, R := i+1, len(nums)-1
		for L < R {
			if num+nums[L]+nums[R] == 0 {
				result = append(result, []int{num, nums[L], nums[R]})
				for L < R && nums[L] == nums[L+1] {
					// L走到下一个和自己不一样的位置
					L++
				}
				for L < R && nums[R] == nums[R-1] {
					// R走到下一个和自己不一样的位置
					R--
				}
				L++
				R--
			} else if num+nums[L]+nums[R] > 0 {
				R--
			} else {
				L++
			}
		}
	}
	return result
}

func main() {

}
