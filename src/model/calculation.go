package model

import "github.com/ssszzzLiiiShuai/go_calculator/src/util"

// 计算公式
type Calculation struct {
	Type         util.CalculationType
	CalculateStr string
	Priority     util.Priority
	Value        float64
	IndexRange   []int
}
