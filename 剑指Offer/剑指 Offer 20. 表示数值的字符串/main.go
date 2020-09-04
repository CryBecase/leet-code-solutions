package main

// https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/

type State int
type CharType int

const (
	StateInitial         State = iota // 起始的空格
	StateIntSign                      // 符号位
	StateInteger                      // 整数部分
	StatePoint                        // 左侧有整数的小数点
	StatePointWithoutInt              // 左侧无整数的小数点
	StateFraction                     // 小数部分
	StateExp                          // 字符 e
	StateExpSign                      // 指数部分的符号位
	StateExpNumber                    // 指数部分的整数部分
	StateEnd                          // 末尾的空格
)

const (
	CharNumber  CharType = iota // 数字
	CharExp                     // 字符 e
	CharPoint                   // 小数点
	CharSign                    // 符号
	CharSpace                   // 空格
	CharIllegal                 // 非法
)

func toCharType(ch byte) CharType {
	switch ch {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return CharNumber
	case 'e', 'E':
		return CharExp
	case '.':
		return CharPoint
	case '+', '-':
		return CharSign
	case ' ':
		return CharSpace
	default:
		return CharIllegal
	}
}

func isNumber(s string) bool {
	// 当前 state 接收 charType 到达下一个 state
	transfer := map[State]map[CharType]State{
		StateInitial: {
			CharSpace:  StateInitial,
			CharNumber: StateInteger,
			CharPoint:  StatePointWithoutInt,
			CharSign:   StateIntSign,
		},
		StateIntSign: {
			CharNumber: StateInteger,
			CharPoint:  StatePointWithoutInt,
		},
		StateInteger: {
			CharNumber: StateInteger,
			CharExp:    StateExp,
			CharPoint:  StatePoint,
			CharSpace:  StateEnd,
		},
		StatePoint: {
			CharNumber: StateFraction,
			CharExp:    StateExp,
			CharSpace:  StateEnd,
		},
		StatePointWithoutInt: {
			CharNumber: StateFraction,
		},
		StateFraction: {
			CharNumber: StateFraction,
			CharExp:    StateExp,
			CharSpace:  StateEnd,
		},
		StateExp: {
			CharNumber: StateExpNumber,
			CharSign:   StateExpSign,
		},
		StateExpSign: {
			CharNumber: StateExpNumber,
		},
		StateExpNumber: {
			CharNumber: StateExpNumber,
			CharSpace:  StateEnd,
		},
		StateEnd: {
			CharSpace: StateEnd,
		},
	}
	state := StateInitial
	for i := 0; i < len(s); i++ {
		typ := toCharType(s[i])
		if _, ok := transfer[state][typ]; !ok {
			return false
		} else {
			state = transfer[state][typ]
		}
	}
	return state == StateInteger || state == StatePoint || state == StateFraction || state == StateExpNumber || state == StateEnd
}

func main() {

}
