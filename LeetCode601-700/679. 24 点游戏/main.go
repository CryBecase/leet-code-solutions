package main

// https://leetcode-cn.com/problems/24-game/

const (
	Target   = 24
	Epsilon  = 1e-6
	Add      = 0
	Multiply = 1
	Subtract = 2
	Divide   = 3
)

func judgePoint24(nums []int) bool {
	var list []float64

	for _, num := range nums {
		list = append(list, float64(num))
	}

	return solve(list)
}

func solve(list []float64) bool {
	if len(list) == 0 {
		return false
	}
	if len(list) == 1 {
		return abs(list[0]-Target) < Epsilon
	}

	size := len(list)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i != j {
				var list2 []float64
				for k := 0; k < size; k++ {
					if k != i && k != j {
						// i j 留着运算之后加进来
						list2 = append(list2, list[k])
					}
				}
				for k := 0; k < 4; k++ {
					if k < 2 && i < j {
						continue
					}
					switch k {
					case Add:
						list2 = append(list2, list[i]+list[j])
					case Multiply:
						list2 = append(list2, list[i]*list[j])
					case Subtract:
						list2 = append(list2, list[i]-list[j])
					case Divide:
						if abs(list[j]) < Epsilon {
							// 被除数是0 跳过
							continue
						} else {
							list2 = append(list2, list[i]/list[j])
						}
					}
					if solve(list2) {
						return true
					}
					list2 = list2[:len(list2)-1]
				}
			}
		}
	}
	return false
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {

}
