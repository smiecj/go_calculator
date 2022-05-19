package calculator

import (
	"fmt"
	"regexp"
)

type compareFunc func(float64, float64) bool

const (
	// 注意 >= 和 <= 优先级要放在前面
	matchComparatorRegex = "(>=|<=|>|<)"
)

var (
	comparatorToMethodMap = map[string]compareFunc{
		">": func(a, b float64) bool {
			return a > b
		},
		"<": func(a, b float64) bool {
			return a < b
		},
		">=": func(a, b float64) bool {
			return a >= b
		},
		"<=": func(a, b float64) bool {
			return a <= b
		},
	}
)

// 比较器入口
// 支持格式: 两个完整浮点数的对比
// 示例: 1.5 >= 0.5
// return: true
// 实现思路: 通过正则表达式校验入参是否规范、比较符和数值之间的分隔
// https://github.com/smiecj/go_calculator/issues/6
func Compare(formula string) (bool, error) {
	// 每次都 compile 可能会有性能问题
	rp, _ := regexp.Compile(matchComparatorRegex)
	indexRangeArr := rp.FindIndex([]byte(formula))

	if nil == indexRangeArr {
		return false, fmt.Errorf(ErrMsgInvalidCalculation)
	}

	// 分别获取分隔符和两边的数值
	comparatorStr := formula[indexRangeArr[0]:indexRangeArr[1]]
	leftValueStr := formula[:indexRangeArr[0]]
	rightValueStr := formula[indexRangeArr[1]:]

	// 数值不合法，也要返回错误
	leftValue, valueErr := Calculate(leftValueStr, nil)
	if nil != valueErr {
		return false, valueErr
	}
	rightValue, valueErr := Calculate(rightValueStr, nil)
	if nil != valueErr {
		return false, valueErr
	}
	// 比较
	return comparatorToMethodMap[comparatorStr](leftValue, rightValue), nil
}
