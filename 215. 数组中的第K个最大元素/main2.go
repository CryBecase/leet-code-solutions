package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 快速选择
func findKthLargest2(nums []int, k int) int {
	if len(nums) < k {
		return -1
	}
	left, right := 0, len(nums)-1
	for left != right {
		rand.Seed(time.Now().UnixNano())
		pivot := left + rand.Intn(right-left)
		nums[right], nums[pivot] = nums[pivot], nums[right]
		p := partition(nums, left, right, k)
		left, right = p[0], p[1]
	}

	return nums[left]
}

func partition(nums []int, L, R, k int) [2]int {
	p0 := L

	less, more := L-1, R
	moreSize, equalSize := 0, 0
	for L < more {
		if nums[L] < nums[R] {
			less++
			nums[less], nums[L] = nums[L], nums[less]
			L++
			//fmt.Printf("小于: %v\n", nums)
		} else if nums[L] > nums[R] {
			more--
			moreSize++
			nums[more], nums[L] = nums[L], nums[more]
			//fmt.Printf("大于: %v\n", nums)
		} else {
			L++
			equalSize++
			//fmt.Printf("等于: %v\n", nums)
		}
	}
	nums[R], nums[more] = nums[more], nums[R]

	if nums[L] == nums[more] {
		equalSize++
	}

	//fmt.Printf("一轮partition: %v less: %d more: %d L: %d R: %d moreSize: %d equalSize: %d\n", nums, less, more, L, R, moreSize, equalSize)
	k = k - (len(nums) - 1 - R)
	if moreSize >= k {
		// 第K大的数在 大于区
		return [2]int{more + 1, R}
	} else if moreSize+equalSize >= k {
		// 在等于区
		return [2]int{less + 1, less + 1}
	} else {
		// 在小于区
		if less < 0 {
			less = 0
		}
		return [2]int{p0, less}
	}
}

func main() {
	test := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest2(test, 4))
}
