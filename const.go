package calculator

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
	ValidAllCharacterRegex = "^([0-9]|[a-z]|[A-Z]|\\(|\\)|\\+|\\-|\\*|\\/).*$"
)

// 运算符
const (
	// operator add
	OperatorAdd = '+'
	// operator sub
	OperatorSub = '-'
	// operator mul
	OperatorMul = '*'
	// operator div
	OperatorDiv = '/'

	// left bracket
	BracketLeft = '('
	// right bracket
	BracketRight = ')'
)

// 优先级
const (
	// highest priority: bracket
	PriorityBracket = 10
	// multiplication and division has medium priority
	PriorityMulAndDiv = 2
	// addition and subtraction has lowest priority
	PriorityAddAndSub = 1
	// some calculation type do not need priority, for example: number
	PriorityEmpty = 0
)

// 一些错误信息
const (
	// err msg: invalid calculation
	ErrMsgInvalidCalculation = "invalid calculation"
	// err msg: some variable do not have declared number
	ErrMsgValueNotExist = "no value to the variable"
	// err msg: div by a zero number
	ErrMsgDivZero = "div 0"

	ErrMsgInvalidComparator = "invalid comparator"
)
