package main

func lengthOfLIS2(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	minTail := make([]int, length)
	minTail[0] = nums[0]
	index := 0 // index+1 就是 结果
	for i := 1; i < length; i++ {
		if nums[i] > minTail[index] {
			index++
			minTail[index] = nums[i]
		} else {
			// 二分查找 找到一个最大的一个数 k 使得 minTail[k] < nums[i] 然后 minTail[k+1] = nums[i]
			L, R, k := 0, index, -1
			for L <= R {
				mid := (L + R) >> 1
				if minTail[mid] < nums[i] {
					L = mid + 1
					k = mid
				} else {
					R = mid - 1
				}
			}
			minTail[k+1] = nums[i]
		}
	}
	return index + 1
}

func main() {
}
