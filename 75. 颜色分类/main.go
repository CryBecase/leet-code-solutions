package main

/**
https://leetcode-cn.com/problems/sort-colors
给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

注意:
不能使用代码库中的排序函数来解决这道题。

示例:

输入: [2,0,2,1,1,0]
输出: [0,0,1,1,2,2]
*/

func sortColors(nums []int) {
	L, R := 0, len(nums)-1
	less, more := L-1, R

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			nums[i], nums[R] = nums[R], nums[i]
			break
		}
	}

	for L < more {
		if nums[L] < nums[R] {
			less++
			nums[less], nums[L] = nums[L], nums[less]
			L++
		} else if nums[L] > nums[R] {
			more--
			nums[more], nums[L] = nums[L], nums[more]
		} else {
			L++
		}
	}
	// 交换
	nums[R], nums[more] = nums[more], nums[R]
}

func main() {

}
