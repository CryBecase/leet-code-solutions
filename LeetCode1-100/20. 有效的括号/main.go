package main

// https://leetcode-cn.com/problems/valid-parentheses/

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]rune, 0)

	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			if c == ')' {
				if stack[len(stack)-1] == '(' {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			} else if c == ']' {
				if stack[len(stack)-1] == '[' {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			} else if c == '}' {
				if stack[len(stack)-1] == '{' {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			}
		}
	}
	return len(stack) == 0
}

func main() {

}
