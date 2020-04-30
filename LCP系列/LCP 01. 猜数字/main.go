package main

// https://leetcode-cn.com/problems/guess-numbers/

func game(guess []int, answer []int) int {
	if len(guess) != len(answer) {
		return -1
	}

	res := 0
	for i := 0; i < len(guess); i++ {
		if guess[i] == answer[i] {
			res++
		}
	}

	return res
}

func main() {

}
