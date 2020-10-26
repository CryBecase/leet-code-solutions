package main

// https://leetcode-cn.com/problems/video-stitching/

func videoStitching(clips [][]int, T int) int {
	helper := make(map[int]int)
	for i, clip := range clips {
		if idx, ok := helper[clip[0]]; ok {
			if clip[1] > clips[idx][1] {
				helper[clip[0]] = i
			}
		} else {
			helper[clip[0]] = i
		}
	}

	curT := 0
	res := 0
	for curT < T {
		preT := curT
		for i := 0; i <= preT; i++ {
			if idx, ok := helper[i]; ok {
				curT = maxFunc(curT, clips[idx][1])
			}
		}
		if preT == curT {
			return -1
		}
		res++
	}

	return res
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
