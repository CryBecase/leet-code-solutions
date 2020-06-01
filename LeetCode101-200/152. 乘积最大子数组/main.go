package main

// https://leetcode-cn.com/problems/maximum-product-subarray/

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	max, min, ans := nums[0], nums[0], nums[0]
	for i, num := range nums {
		if i == 0 {
			continue
		}
		lastMax, lastMin := max, min
		max = maxFunc(maxFunc(num, num*lastMax), num*lastMin)
		min = minFunc(minFunc(num, num*lastMax), num*lastMin)
		ans = maxFunc(max, ans)
	}
	return ans
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func minFunc(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {

}
