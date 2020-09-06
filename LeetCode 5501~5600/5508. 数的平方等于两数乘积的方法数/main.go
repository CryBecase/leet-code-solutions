package main

func numTriplets(nums1 []int, nums2 []int) int {
	m2 := make(map[int]int)
	for i := 0; i < len(nums2)-1; i++ {
		for j := i + 1; j < len(nums2); j++ {
			m2[nums2[i]*nums2[j]]++
		}
	}
	m1 := make(map[int]int)
	for i := 0; i < len(nums1)-1; i++ {
		for j := i + 1; j < len(nums1); j++ {
			m1[nums1[i]*nums1[j]]++
		}
	}

	return handle(nums1, m2) + handle(nums2, m1)
}

func handle(nums1 []int, m map[int]int) int {
	res := 0
	for _, n := range nums1 {
		res += m[n*n]
	}
	return res
}

func main() {

}
