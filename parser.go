package calculator

// 判断字符类型
// 当前: 确认 -10 这种应该当作是数字还是计算符？ -- 应该当作运算符，然后左边的数当作0
// 当前: 不能接运算符的规则改一下，可以接+,-，不能接 * /
func GetCharacterCalculationType(c rune) CalculationType {
	if OperatorAdd == c || OperatorSub == c || OperatorMul == c || OperatorDiv == c {
		return Operator
	} else if c >= '0' && c <= '9' {
		return Number
	} else if BracketLeft == c || BracketRight == c {
		return Bracket
	} else {
		return Invalid
	}
}

// 判断变量类型：常数OR变量
func GetVariableCalculationType(v string) CalculationType {
	if 0 == len(v) {
		return Invalid
	} else if v[0] >= '0' && v[0] <= '9' {
		return Number
	} else {
		return Variable
	}
}

// 获取运算符的优先级
func GetCalculatorPriority(c rune) Priority {
	if OperatorAdd == c || OperatorSub == c {
		return PriorityAddAndSub
	} else {
		return PriorityMulAndDiv
	}
}

// 公共方法: 判断运算符是否是 * /
func IsMulOrDivOrCalculation(c rune) bool {
	return OperatorMul == c || OperatorDiv == c
}
