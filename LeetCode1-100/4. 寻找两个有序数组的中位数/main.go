package main

/**
https://leetcode-cn.com/problems/median-of-two-sorted-arrays/

给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。
*/

/**
个人做题时自拟的例子
a{1, 4, 7, 16, 17, 21, 23}
b{2, 8, 14, 15, 26}
k = 7
*/

/**
求第K个数解法 找到 i 和 j 满足以下条件
条件一: i + j = k
条件二: a[i-1] <= b[j] && b[j-1] <= a[i] (保证了这 i + j 个数连续)
第K个数 = max(a[i-1], b[j-1])

在a[]中二分查找 i
	0 <= i <= n
	j = k-i, 0 <= j <= m
得 k-m <= i <= k
得 max(0, k-m) <= i <= min(n, k)
*/
func findMedianSortedArrays(a []int, b []int) float64 {
	n, m := len(a), len(b)
	if n == 0 {
		return handleSingleArray(b)
	}
	if m == 0 {
		return handleSingleArray(a)
	}
	totalLength := n + m
	if totalLength%2 == 0 {
		return (findKthNum(a, b, (n+m)/2+1) + findKthNum(a, b, (n+m)/2)) / 2
	} else {
		return findKthNum(a, b, (n+m)/2+1)
	}
}

func findKthNum(a, b []int, k int) float64 {
	n, m := len(a), len(b)
	iL, iR := max(0, k-m), min(n, k)
	for iL < iR {
		i := (iL + iR) / 2
		j := k - i
		if b[j-1] > a[i] {
			// b[j] 肯定大于 a[i-1] 所以 i 左边全部不考虑
			iL = i + 1
		} else {
			iR = i
		}
	}
	i := iL
	j := k - i
	if i == 0 {
		return float64(b[j-1])
	}
	if i == k {
		return float64(a[i-1])
	}
	return float64(max(a[i-1], b[j-1]))
}

func handleSingleArray(arr []int) float64 {
	length := len(arr)
	if length%2 == 0 {
		return float64(arr[length/2]+arr[length/2-1]) / 2
	} else {
		return float64(arr[length/2])
	}
}

func max(n, m int) int {
	if n >= m {
		return n
	} else {
		return m
	}
}

func min(n, m int) int {
	if n <= m {
		return n
	} else {
		return m
	}
}

func main() {

}
