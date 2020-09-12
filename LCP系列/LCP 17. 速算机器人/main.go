package main

// https://leetcode-cn.com/problems/nGK0Fy/

var (
	x = 1
	y = 0
)

func calculate(s string) int {
	x = 1
	y = 0
	for _, c := range s {
		if c == 'A' {
			A()
		}
		if c == 'B' {
			B()
		}
	}

	return x + y
}

func A() {
	x = 2*x + y
}

func B() {
	y = 2*y + x
}

func main() {

}
