package main

// https://leetcode-cn.com/problems/sort-colors

func sortColors(nums []int) {
	p1, p2 := 0, len(nums)-1
	for p := 0; p <= p2; p++ {
		for ; p <= p2 && nums[p] == 2; p2-- {
			nums[p], nums[p2] = nums[p2], nums[p]
		}
		if nums[p] == 0 {
			nums[p], nums[p1] = nums[p1], nums[p]
			p1++
		}
	}
}

func main() {

}
