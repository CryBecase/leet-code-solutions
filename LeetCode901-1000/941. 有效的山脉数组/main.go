package main

// https://leetcode-cn.com/problems/valid-mountain-array/

func validMountainArray(a []int) bool {
	i, n := 0, len(a)

	// 递增扫描
	for ; i+1 < n && a[i] < a[i+1]; i++ {
	}

	// 最高点不能是数组的第一个位置或最后一个位置
	if i == 0 || i == n-1 {
		return false
	}

	// 递减扫描
	for ; i+1 < n && a[i] > a[i+1]; i++ {
	}

	return i == n-1
}

func main() {

}
