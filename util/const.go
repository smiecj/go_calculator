// package util is some constant and type which is use for calculation.
package util

// calculation type in the formula
type CalculationType uint

// calculate priority, for example multiplication has higher priority than addition
type Priority uint

const (
	// operator such as *
	Operator CalculationType = 0
	// explicit number
	Number CalculationType = 1
	// bracket: ( and )
	Bracket CalculationType = 2
	// such as v1 and v2 which value is declared by user
	Variable CalculationType = 3
	// invalid input
	Invalid CalculationType = 4
)

// 所有合法字符通配符
const (
	// invalid formular check regex
	VALID_ALL_CHARACTER_REGEX = "^([0-9]|[a-z]|[A-Z]|\\(|\\)|\\+|\\-|\\*|\\/).*$"
)

// 运算符
const (
	// operator add
	OPERATOR_ADD = '+'
	// operator sub
	OPERATOR_SUB = '-'
	// operator mul
	OPERATOR_MUL = '*'
	// operator div
	OPERATOR_DIV = '/'

	// left bracket
	BRACKET_LEFT = '('
	// right bracket
	BRACKET_RIGHT = ')'
)

// 优先级
const (
	// highest priority: bracket
	PRIORITY_BRACKET = 10
	// multiplication and division has medium priority
	PRIORITY_MULANDDIV = 2
	// addition and subtraction has lowest priority
	PRIORITY_ADDANDSUB = 1
	// some calculation type do not need priority, for example: number
	PRIORITY_EMPTY = 0
)

// 一些错误信息
const (
	// err msg: invalid calculation
	ERR_MSG_INVALID_CALCULATION = "不合法的运算符！"
	// err msg: some variable do not have declared number
	ERR_MSG_VALUE_NOT_EXIST = "参与计算的变量没有实际值！"
	// err msg: div by a zero number
	ERR_MSG_DIV_ZERO = "除0错误！"
)
