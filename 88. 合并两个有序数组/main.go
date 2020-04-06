package main

import "fmt"

/**
https://leetcode-cn.com/problems/merge-sorted-array
给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。

说明:

初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
示例:

输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

输出: [1,2,2,3,5,6]
*/

// 从前向后
func merge(nums1 []int, m int, nums2 []int, n int) {
	helper := make([]int, 0, m+n)
	m2, n2 := 0, 0
	for m2 < m && n2 < n {
		if nums1[m2] <= nums2[n2] {
			helper = append(helper, nums1[m2])
			m2++
		} else {
			helper = append(helper, nums2[n2])
			n2++
		}
	}
	for m2 < m {
		helper = append(helper, nums1[m2])
		m2++
	}
	for n2 < n {
		helper = append(helper, nums2[n2])
		n2++
	}
	copy(nums1, helper)
}

// 从后往前
func merge2(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, p := m-1, n-1, m+n-1
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] <= nums2[p2] {
			nums1[p] = nums2[p2]
			p--
			p2--
		} else {
			nums1[p] = nums1[p1]
			p--
			p1--
		}
	}
	for p1 >= 0 {
		nums1[p] = nums1[p1]
		p--
		p1--
	}
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p--
		p2--
	}
}
func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge2(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}
