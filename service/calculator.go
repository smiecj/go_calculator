package service

import (
	"fmt"
	"github.com/smiecj/go_calculator/model"
	"github.com/smiecj/go_calculator/util"
	"regexp"
	"strconv"
	"strings"
)

var (
	characterInvalidRegex = regexp.MustCompile(util.VALID_ALL_CHARACTER_REGEX)
)

// 判断整个表达式是否符合规范，并返回小括号对应的下标
func isValidCalculation(calculation string) (map[int]int, error) {
	bracketIndexMap := make(map[int]int)
	bracketLeftIndexArr := make([]int, 0)
	if 0 == len(calculation) {
		//log.Printf("[isValidCalculation] 计算公式长度为0，不合法!")
		return bracketIndexMap, fmt.Errorf(util.ERR_MSG_INVALID_CALCULATION)
	}

	// 有非法字符
	if !characterInvalidRegex.MatchString(calculation) {
		return bracketIndexMap, fmt.Errorf(util.ERR_MSG_INVALID_CALCULATION)
	}

	for index, character := range calculation {
		// 判断左括号的合法性：不能立刻接运算符
		if util.BRACKET_LEFT == character {
			if index+1 < len(calculation) && util.Operator == util.GetCharacterCalculationType(character) {
				return bracketIndexMap, fmt.Errorf(util.ERR_MSG_INVALID_CALCULATION)
			}
			bracketLeftIndexArr = append(bracketLeftIndexArr, index)
		} else if util.BRACKET_RIGHT == character {
			// 没有匹配的左括号，不合法
			if 0 == len(bracketLeftIndexArr) {
				return bracketIndexMap, fmt.Errorf(util.ERR_MSG_INVALID_CALCULATION)
			}
			// 右括号左边是运算符，也不合法
			if util.Operator == util.GetCharacterCalculationType(rune(calculation[index-1])) {
				return bracketIndexMap, fmt.Errorf(util.ERR_MSG_INVALID_CALCULATION)
			}
			// pop
			bracketIndexMap[bracketLeftIndexArr[len(bracketLeftIndexArr)-1]] = index
			bracketLeftIndexArr = bracketLeftIndexArr[:len(bracketLeftIndexArr)-1]
		} else if util.Operator == util.GetCharacterCalculationType(character) {
			// 运算符前后不能也是运算符，必须有数字
			if index == len(calculation)-1 {
				return bracketIndexMap, fmt.Errorf(util.ERR_MSG_INVALID_CALCULATION)
			}
			if index == 0 {
				return bracketIndexMap, fmt.Errorf(util.ERR_MSG_INVALID_CALCULATION)
			}
			if index < len(calculation)-1 &&
				util.Operator == util.GetCharacterCalculationType(rune(calculation[index+1])) {
				return bracketIndexMap, fmt.Errorf(util.ERR_MSG_INVALID_CALCULATION)
			}
		}
	}

	// 最后递归完成之后，还要判断是否所有的左括号都和右括号匹配成功了
	if 0 == len(bracketLeftIndexArr) {
		return bracketIndexMap, nil
	}
	return bracketIndexMap, fmt.Errorf(util.ERR_MSG_INVALID_CALCULATION)
}

// 计算递归方法，主要逻辑
func recursiveCalculate(startIndex, endIndex int, calculation string, bracketIndexMap map[int]int,
	valueMap map[string]float64) (float64, error) {
	if startIndex >= endIndex {
		return 0, nil
	}
	// 如果当前运算本身最外层就是被括号包括，直接去掉后递归计算
	if util.BRACKET_LEFT == calculation[startIndex] && util.BRACKET_RIGHT == calculation[endIndex] {
		return recursiveCalculate(startIndex+1, endIndex-1, calculation, bracketIndexMap, valueMap)
	}

	// 遍历整个计算公式，并生成计算公式数组
	calculationArr := make([]*model.Calculation, 0)
	currentIndex, lastStartIndex := startIndex, startIndex
	for currentIndex <= endIndex {
		if util.BRACKET_LEFT == calculation[currentIndex] {
			bracketEndIndex := bracketIndexMap[currentIndex]
			currentCalculation := model.Calculation{
				Type:         util.Bracket,
				CalculateStr: calculation[currentIndex+1 : bracketEndIndex],
				Priority:     util.PRIORITY_BRACKET,
				IndexRange:   []int{currentIndex + 1, bracketEndIndex - 1},
			}
			calculationArr = append(calculationArr, &currentCalculation)
			lastStartIndex = bracketEndIndex + 1
			currentIndex = lastStartIndex
		} else if util.Operator == util.GetCharacterCalculationType(rune(calculation[currentIndex])) {
			operatorPriority := util.GetCalculatorPriority(rune(calculation[currentIndex]))
			// 添加运算符前面的变量/数字
			if currentIndex > lastStartIndex {
				lastVariable := calculation[lastStartIndex:currentIndex]
				currentVariable := model.Calculation{
					Type:         util.GetVariableCalculationType(lastVariable),
					CalculateStr: lastVariable,
					Priority:     util.PRIORITY_EMPTY,
				}
				calculationArr = append(calculationArr, &currentVariable)
			}
			currentCalculation := model.Calculation{
				Type:         util.Operator,
				CalculateStr: string(calculation[currentIndex]),
				Priority:     operatorPriority,
			}
			calculationArr = append(calculationArr, &currentCalculation)
			currentIndex++
			lastStartIndex = currentIndex
		} else if currentIndex == endIndex {
			// 添加最后一个变量
			lastVariable := calculation[lastStartIndex : endIndex+1]
			currentVariable := model.Calculation{
				Type:         util.GetVariableCalculationType(lastVariable),
				CalculateStr: lastVariable,
				Priority:     util.PRIORITY_EMPTY,
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
		if util.Bracket == currentCalculation.Type {
			bracketValue, err := recursiveCalculate(currentCalculation.IndexRange[0], currentCalculation.IndexRange[1],
				calculation, bracketIndexMap, valueMap)
			if nil != err {
				return 0, err
			}
			currentCalculation.Type, currentCalculation.Value, currentCalculation.Priority =
				util.Number, bracketValue, util.PRIORITY_EMPTY
		} else if util.Variable == currentCalculation.Type {
			value, ok := valueMap[currentCalculation.CalculateStr]
			if !ok {
				return 0, fmt.Errorf(util.ERR_MSG_VALUE_NOT_EXIST)
			}
			currentCalculation.Type, currentCalculation.Value = util.Number, value
		} else if util.Number == currentCalculation.Type {
			value, err := strconv.ParseFloat(currentCalculation.CalculateStr, 64)
			if nil != err {
				return 0, err
			}
			currentCalculation.Type, currentCalculation.Value = util.Number, value
		}
	}

	// 优先级: * /
	tempCalculationArr := make([]*model.Calculation, 0)
	for index := 0; index < len(calculationArr); index++ {
		currentCalculation := calculationArr[index]
		if util.PRIORITY_MULANDDIV != currentCalculation.Priority {
			tempCalculationArr = append(tempCalculationArr, currentCalculation)
		} else {
			lastCalculation, nextCalculation := tempCalculationArr[len(tempCalculationArr)-1],
				calculationArr[index+1]
			newCalculation := model.Calculation{
				Type:     util.Number,
				Priority: util.PRIORITY_EMPTY,
			}
			if util.OPERATOR_MUL == currentCalculation.CalculateStr[0] {
				newCalculation.Value = lastCalculation.Value * nextCalculation.Value
			} else {
				if 0 == nextCalculation.Value {
					return 0, fmt.Errorf(util.ERR_MSG_DIV_ZERO)
				}
				newCalculation.Value = lastCalculation.Value / nextCalculation.Value
			}
			// 注意由于temp 数组放了上一个对象，这里要pop
			tempCalculationArr = tempCalculationArr[:len(tempCalculationArr)-1]
			tempCalculationArr = append(tempCalculationArr, &newCalculation)
			index++
		}
	}

	// 优先级: + -
	calculationArr = tempCalculationArr
	retValue := calculationArr[0].Value
	for index := 1; index < len(calculationArr); index++ {
		currentCalculation := calculationArr[index]
		if util.PRIORITY_ADDANDSUB != currentCalculation.Priority {
			continue
		}
		nextCalculation := calculationArr[index+1]
		if util.OPERATOR_ADD == currentCalculation.CalculateStr[0] {
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
