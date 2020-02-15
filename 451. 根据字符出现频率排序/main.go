package main

import "fmt"

/**
https://leetcode-cn.com/problems/sort-characters-by-frequency
给定一个字符串，请将字符串里的字符按照出现的频率降序排列。

示例 1:

输入:
"tree"

输出:
"eert"

解释:
'e'出现两次，'r'和't'都只出现一次。
因此'e'必须出现在'r'和't'之前。此外，"eetr"也是一个有效的答案。
示例 2:

输入:
"cccaaa"

输出:
"cccaaa"

解释:
'c'和'a'都出现三次。此外，"aaaccc"也是有效的答案。
注意"cacaca"是不正确的，因为相同的字母必须放在一起。
示例 3:

输入:
"Aabb"

输出:
"bbAa"

解释:
此外，"bbaA"也是一个有效的答案，但"Aabb"是不正确的。
注意'A'和'a'被认为是两种不同的字符。
*/

func frequencySort(s string) string {
	// string to bytes
	sByte := []byte(s)
	// 统计byte出现的频率
	helper := make(map[byte]int)
	for i := range sByte {
		helper[sByte[i]]++
	}
	// 初始化桶
	max := 0
	for i := range helper {
		max = maxFunc(max, helper[i])
	}
	// 一维数组的下标是频率 一维数组里的数组是频率相同的byte
	bucket := make([][]byte, max+1, max+1)
	for key := range helper {
		frequent := helper[key]
		bucket[frequent] = append(bucket[frequent], key)
	}
	// 从桶中拿出byte
	var result []byte
	for i := len(bucket) - 1; i >= 0; i-- {
		for j := 0; j < len(bucket[i]); j++ {
			for m := i; m > 0; m-- {
				result = append(result, bucket[i][j])
			}
		}
	}

	return string(result)
}

func maxFunc(i int, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func main() {
	fmt.Println(frequencySort("tree"))
}
