// declare calculation model
package calculator

// 计算公式
type Calculation struct {
	Type         CalculationType
	CalculateStr string
	Priority     Priority
	Value        float64
	IndexRange   []int
}
