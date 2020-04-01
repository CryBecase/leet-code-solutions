package main

// https://leetcode-cn.com/problems/valid-parentheses/

type stack struct {
	data     []int32
	length   int
	capacity int
}

func InitStack(l int) *stack {
	return &stack{
		data:     make([]int32, 0, l),
		length:   0,
		capacity: l,
	}
}

func (s *stack) Push(data int32) {
	if s.length+1 > s.capacity {
		s.capacity <<= 1
		d := s.data
		s.data = make([]int32, 0, s.capacity)
		copy(s.data, d)
	}
	s.data = append(s.data, data)
	s.length++
}

func (s *stack) Pop() int32 {
	if s.length <= 0 {
		return 0
	}
	t := s.data[s.length-1]
	s.data = s.data[:s.length-1]
	s.length--
	return t
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := InitStack(len(s))
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack.Push(c)
		} else {
			pop := stack.Pop()
			if c == ')' && pop != '(' {
				return false
			} else if c == '}' && pop != '{' {
				return false
			} else if c == ']' && pop != '[' {
				return false
			}
		}
	}
	if stack.length > 0 {
		return false
	}
	return true
}

func main() {

}
