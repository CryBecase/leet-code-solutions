package main

// https://leetcode-cn.com/problems/decompress-run-length-encoded-list/

func decompressRLElist(nums []int) []int {
	if len(nums)%2 != 0 {
		return nil
	}

	i := 0
	res := make([]int, 0)
	for i < len(nums) {
		frep, val := nums[i], nums[i+1]
		for j := 0; j < frep; j++ {
			res = append(res, val)
		}
		i = i + 2
	}

	return res
}

func main() {

}
