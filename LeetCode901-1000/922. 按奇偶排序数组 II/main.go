package main

// https://leetcode-cn.com/problems/sort-array-by-parity-ii/

func sortArrayByParityII(A []int) []int {
	arr1 := make([]int, 0)
	arr2 := make([]int, 0)
	for _, n := range A {
		if n%2 == 0 {
			arr1 = append(arr1, n)
		} else {
			arr2 = append(arr2, n)
		}
	}

	res := make([]int, 0)
	p1, p2 := 0, 0
	for p1 < len(arr1) && p2 < len(arr2) {
		res = append(res, arr1[p1])
		res = append(res, arr2[p2])
		p1++
		p2++
	}
	for p1 < len(arr1) {
		res = append(res, arr1[p1])
		p1++
	}
	for p2 < len(arr2) {
		res = append(res, arr2[p2])
		p2++
	}
	return res
}

func main() {

}
