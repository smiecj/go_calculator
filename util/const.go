package util

type CalculationType uint
type Priority uint

const (
	Operator CalculationType = 0
	Number   CalculationType = 1
	Bracket  CalculationType = 2
	Variable CalculationType = 3
	Invalid  CalculationType = 4
)

// 所有合法字符通配符
const (
	VALID_ALL_CHARACTER_REGEX = "^([0-9]|[a-z]|[A-Z]|\\(|\\)|\\+|\\-|\\*|\\/).*$"
)

// 运算符
const (
	OPERATOR_ADD = '+'
	OPERATOR_SUB = '-'
	OPERATOR_MUL = '*'
	OPERATOR_DIV = '/'

	BRACKET_LEFT  = '('
	BRACKET_RIGHT = ')'
)

// 优先级
const (
	PRIORITY_BRACKET   = 10
	PRIORITY_MULANDDIV = 2
	PRIORITY_ADDANDSUB = 1
	PRIORITY_EMPTY     = 0
)

const (
	ERR_MSG_INVALID_CALCULATION = "不合法的运算符！"
	ERR_MSG_VALUE_NOT_EXIST     = "参与计算的变量没有实际值！"
	ERR_MSG_DIV_ZERO            = "除0错误！"
)
