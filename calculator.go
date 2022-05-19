// package service declare the main calculate logic
package calculator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	characterInvalidRegex = regexp.MustCompile(ValidAllCharacterRegex)
)

// 判断整个表达式是否符合规范，并返回小括号对应的下标
func isValidCalculation(calculation string) (map[int]int, error) {
	bracketIndexMap := make(map[int]int)
	bracketLeftIndexArr := make([]int, 0)
	if 0 == len(calculation) {
		return bracketIndexMap, fmt.Errorf(ErrMsgInvalidCalculation)
	}

	// 有非法字符
	if !characterInvalidRegex.MatchString(calculation) {
		return bracketIndexMap, fmt.Errorf(ErrMsgInvalidCalculation)
	}

	for index, character := range calculation {
		// 判断左括号的合法性：不能立刻接乘、除运算符
		if BracketLeft == character {
			if index+1 < len(calculation) && IsMulOrDivOrCalculation(character) {
				return bracketIndexMap, fmt.Errorf(ErrMsgInvalidCalculation)
			}
			bracketLeftIndexArr = append(bracketLeftIndexArr, index)
		} else if BracketRight == character {
			// 没有匹配的左括号，不合法
			if 0 == len(bracketLeftIndexArr) {
				return bracketIndexMap, fmt.Errorf(ErrMsgInvalidCalculation)
			}
			// 右括号左边是运算符，不合法
			if Operator == GetCharacterCalculationType(rune(calculation[index-1])) {
				return bracketIndexMap, fmt.Errorf(ErrMsgInvalidCalculation)
			}
			// pop
			bracketIndexMap[bracketLeftIndexArr[len(bracketLeftIndexArr)-1]] = index
			bracketLeftIndexArr = bracketLeftIndexArr[:len(bracketLeftIndexArr)-1]
		} else if Operator == GetCharacterCalculationType(character) {
			// 运算符前后不能也是运算符，必须有数字
			if index == len(calculation)-1 {
				return bracketIndexMap, fmt.Errorf(ErrMsgInvalidCalculation)
			}
			if index == 0 && IsMulOrDivOrCalculation(character) {
				return bracketIndexMap, fmt.Errorf(ErrMsgInvalidCalculation)
			}
			if index < len(calculation)-1 &&
				Operator == GetCharacterCalculationType(rune(calculation[index+1])) {
				return bracketIndexMap, fmt.Errorf(ErrMsgInvalidCalculation)
			}
		}
	}

	// 最后递归完成之后，还要判断是否所有的左括号都和右括号匹配成功了
	if 0 == len(bracketLeftIndexArr) {
		return bracketIndexMap, nil
	}
	return bracketIndexMap, fmt.Errorf(ErrMsgInvalidCalculation)
}

// 计算递归方法，主要逻辑
func recursiveCalculate(startIndex, endIndex int, calculation string, bracketIndexMap map[int]int,
	valueMap map[string]float64) (float64, error) {
	if startIndex > endIndex {
		return 0, nil
	}
	// 如果当前运算本身最外层就是被括号包括，直接去掉后递归计算
	// 20201227: fix-1: 不能以两头分别是左括号和右括号就认为 是同一对括号，还是要用 bracketIndexMap 做位置判断，必须是对应的括号
	if BracketLeft == calculation[startIndex] && endIndex == bracketIndexMap[startIndex] {
		return recursiveCalculate(startIndex+1, endIndex-1, calculation, bracketIndexMap, valueMap)
	}

	// 遍历整个计算公式，并生成计算公式数组
	calculationArr := make([]*Calculation, 0)
	currentIndex, lastStartIndex := startIndex, startIndex
	for currentIndex <= endIndex {
		if BracketLeft == calculation[currentIndex] {
			bracketEndIndex := bracketIndexMap[currentIndex]
			currentCalculation := Calculation{
				Type:         Bracket,
				CalculateStr: calculation[currentIndex+1 : bracketEndIndex],
				Priority:     PriorityBracket,
				IndexRange:   []int{currentIndex + 1, bracketEndIndex - 1},
			}
			calculationArr = append(calculationArr, &currentCalculation)
			lastStartIndex = bracketEndIndex + 1
			currentIndex = lastStartIndex
		} else if Operator == GetCharacterCalculationType(rune(calculation[currentIndex])) {
			operatorPriority := GetCalculatorPriority(rune(calculation[currentIndex]))
			// 添加运算符前面的变量/数字
			if currentIndex > lastStartIndex {
				lastVariable := calculation[lastStartIndex:currentIndex]
				currentVariable := Calculation{
					Type:         GetVariableCalculationType(lastVariable),
					CalculateStr: lastVariable,
					Priority:     PriorityEmpty,
				}
				calculationArr = append(calculationArr, &currentVariable)
			}
			currentCalculation := Calculation{
				Type:         Operator,
				CalculateStr: string(calculation[currentIndex]),
				Priority:     operatorPriority,
			}
			calculationArr = append(calculationArr, &currentCalculation)
			currentIndex++
			lastStartIndex = currentIndex
		} else if currentIndex == endIndex {
			// 添加最后一个变量
			lastVariable := calculation[lastStartIndex : endIndex+1]
			currentVariable := Calculation{
				Type:         GetVariableCalculationType(lastVariable),
				CalculateStr: lastVariable,
				Priority:     PriorityEmpty,
			}
			calculationArr = append(calculationArr, &currentVariable)
			currentIndex++
			lastStartIndex = currentIndex
		} else {
			currentIndex++
		}
	}

	// 开始计算
	// 首先计算优先级最高的括号运算和变量替换。括号运算由于走递归，最后一定能替换成数值
	// 变量替换和括号替换放到一起的原因，是可以区分括号计算后的数值（直接是float）和数值字符串转float的不同
	for _, currentCalculation := range calculationArr {
		if Bracket == currentCalculation.Type {
			bracketValue, err := recursiveCalculate(currentCalculation.IndexRange[0], currentCalculation.IndexRange[1],
				calculation, bracketIndexMap, valueMap)
			if nil != err {
				return 0, err
			}
			currentCalculation.Type, currentCalculation.Value, currentCalculation.Priority =
				Number, bracketValue, PriorityEmpty
		} else if Variable == currentCalculation.Type {
			value, ok := valueMap[currentCalculation.CalculateStr]
			if !ok {
				return 0, fmt.Errorf(ErrMsgValueNotExist)
			}
			currentCalculation.Type, currentCalculation.Value = Number, value
		} else if Number == currentCalculation.Type {
			value, err := strconv.ParseFloat(currentCalculation.CalculateStr, 64)
			if nil != err {
				return 0, err
			}
			currentCalculation.Type, currentCalculation.Value = Number, value
		}
	}

	// 优先级: 负数
	// 判断方式: 负号前面不是数字
	// todo: CalculateStr[0] 类似这种写法不要出现，不够直观，通过对象方法返回操作符
	tempCalculationArr := make([]*Calculation, 0)
	for index := 0; index < len(calculationArr); index++ {
		if calculationArr[index].Type == Operator && calculationArr[index].CalculateStr[0] == OperatorSub &&
			(index == 0 || calculationArr[index-1].Type != Number) {
			// 将下一个整数整合成一个负数
			calculationArr[index+1].Value = -calculationArr[index+1].Value
			tempCalculationArr = append(tempCalculationArr, calculationArr[index+1])
			index++
		} else {
			tempCalculationArr = append(tempCalculationArr, calculationArr[index])
		}
	}
	calculationArr = tempCalculationArr

	// 优先级: * /
	tempCalculationArr = make([]*Calculation, 0)
	for index := 0; index < len(calculationArr); index++ {
		currentCalculation := calculationArr[index]
		if PriorityMulAndDiv != currentCalculation.Priority {
			tempCalculationArr = append(tempCalculationArr, currentCalculation)
		} else {
			lastCalculation, nextCalculation := tempCalculationArr[len(tempCalculationArr)-1],
				calculationArr[index+1]
			newCalculation := Calculation{
				Type:     Number,
				Priority: PriorityEmpty,
			}
			if OperatorMul == currentCalculation.CalculateStr[0] {
				newCalculation.Value = lastCalculation.Value * nextCalculation.Value
			} else {
				if 0 == nextCalculation.Value {
					return 0, fmt.Errorf(ErrMsgDivZero)
				}
				newCalculation.Value = lastCalculation.Value / nextCalculation.Value
			}
			// 注意由于temp 数组放了上一个对象，这里要pop
			tempCalculationArr = tempCalculationArr[:len(tempCalculationArr)-1]
			tempCalculationArr = append(tempCalculationArr, &newCalculation)
			index++
		}
	}
	calculationArr = tempCalculationArr

	// 优先级: + -
	var retValue float64 = 0
	if len(calculationArr) > 0 {
		retValue = calculationArr[0].Value
	}
	for index := 1; index < len(calculationArr); index++ {
		currentCalculation := calculationArr[index]
		if PriorityAddAndSub != currentCalculation.Priority {
			continue
		}
		nextCalculation := calculationArr[index+1]
		if OperatorAdd == currentCalculation.CalculateStr[0] {
			retValue += nextCalculation.Value
		} else {
			retValue -= nextCalculation.Value
		}
		index++
	}
	return retValue, nil
}

// 计算主方法入口
func Calculate(calculation string, valValueMap map[string]float64) (float64, error) {
	calculation = strings.ReplaceAll(calculation, " ", "")
	bracketIndexMap, err := isValidCalculation(calculation)
	if nil != err {
		return 0, err
	} else {
		return recursiveCalculate(0, len(calculation)-1, calculation, bracketIndexMap, valValueMap)
	}
}
