package main

// https://leetcode-cn.com/problems/largest-rectangle-in-histogram/

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	if len(heights) == 1 {
		return heights[0]
	}
	minIdx := 0
	for i, h := range heights {
		if h < heights[minIdx] {
			minIdx = i
		}
	}
	leftArea := largestRectangleArea(heights[:minIdx])
	rightArea := 0
	if minIdx+1 < len(heights) {
		rightArea = largestRectangleArea(heights[minIdx+1:])
	}
	maxArea := maxFunc(leftArea, rightArea)

	return maxFunc(maxArea, heights[minIdx]*len(heights))
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
