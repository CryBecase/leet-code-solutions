package main

import "fmt"

// https://leetcode-cn.com/problems/sum-of-mutated-array-closest-to-target/

func findBestValue(arr []int, target int) int {
	avg := target / len(arr)
	bigArr := make([]int, 0)
	nextTarget := target
	max := 0
	for _, num := range arr {
		if num > avg {
			bigArr = append(bigArr, num)
		} else {
			nextTarget -= num
		}
		if num > max {
			max = num
		}
	}
	// 全是比平均值大的 取平均值
	if len(arr) == len(bigArr) {
		if abs(avg*len(arr)-target) <= abs((avg+1)*len(arr)-target) {
			return avg
		}
		return avg + 1
	}
	if len(bigArr) == 0 {
		return max
	}
	return findBestValue(bigArr, nextTarget)
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}

func main() {
	value := findBestValue([]int{2, 3, 5}, 10)
	fmt.Println(value)
}
