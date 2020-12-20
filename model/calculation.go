// package model declare calculation model
package model

import "github.com/smiecj/go_calculator/util"

// 计算公式
type Calculation struct {
	Type         util.CalculationType
	CalculateStr string
	Priority     util.Priority
	Value        float64
	IndexRange   []int
}
