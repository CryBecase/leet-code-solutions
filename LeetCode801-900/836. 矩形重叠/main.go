package main

// https://leetcode-cn.com/problems/rectangle-overlap/

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// 排除不在的情况
	return !(rec1[2] <= rec2[0] || // left
		rec1[3] <= rec2[1] || // bottom
		rec1[0] >= rec2[2] || // right
		rec1[1] >= rec2[3]) // top
}

func main() {

}
