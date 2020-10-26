package util

// 判断字符类型
func GetCharacterCalculationType(c rune) CalculationType {
	if OPERATOR_ADD == c || OPERATOR_SUB == c || OPERATOR_MUL == c || OPERATOR_DIV == c {
		return Operator
	} else if c >= '0' && c <= '9' {
		return Number
	} else if BRACKET_LEFT == c || BRACKET_RIGHT == c {
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
	if OPERATOR_ADD == c || OPERATOR_SUB == c {
		return PRIORITY_ADDANDSUB
	} else {
		return PRIORITY_MULANDDIV
	}
}
