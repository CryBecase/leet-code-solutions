package main

import (
	"fmt"
	"math"
)

// https://leetcode-cn.com/problems/string-to-integer-atoi/

//           |    ' '      +/-      number     other
//   start   |   start     type    in_number    end
//   type    |    end      end     in_number    end
// in_number |    end      end     in_number    end
//   end     |    end      end       end        end

func myAtoi(str string) int {
	state := "start"
	t := '+'
	table := map[string][]string{
		"start":     {"start", "type", "in_number", "end"},
		"type":      {"end", "end", "in_number", "end"},
		"in_number": {"end", "end", "in_number", "end"},
		"end":       {"end", "end", "end", "end"},
	}
	result := 0
	for _, c := range str {
		if c == ' ' {
			state = table[state][0]
		} else if c == '+' || c == '-' {
			state = table[state][1]
		} else if c >= '0' && c <= '9' {
			state = table[state][2]
		} else {
			state = table[state][3]
		}

		if state == "end" {
			if t == '-' {
				return -result
			}
			return result
		} else if state == "in_number" {
			if t == '+' {
				if math.MaxInt32/10 < result {
					return math.MaxInt32
				} else if math.MaxInt32/10 == result && int(c-'0') > 7 {
					return math.MaxInt32
				}
			} else {
				if math.MaxInt32/10 < result {
					return math.MinInt32
				} else if math.MaxInt32/10 == result && int(c-'0') > 8 {
					return math.MinInt32
				}
			}

			result = result*10 + int(c-'0')
		} else if state == "type" {
			t = c
		}
	}
	if t == '-' {
		return -result
	}
	return result
}

func main() {
	fmt.Println(math.MinInt32)
}
