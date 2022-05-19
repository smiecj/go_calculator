# go_calculator
使用 go 写的计算器

使用方式：可参考test/calculate_test.go 中的代码

第一种方式：纯数值计算
如: 3 * (1 + 2)

第二种方式：数值+变量计算
v1 * (2 + v3)

注意第二种方式需要保证每一个变量的值都在 valueMap 当中有，否则最后Calculate 方法会返回error

## 使用方法

### 计算器
```go
// 不带变量
import calculator "github.com/smiecj/go_calculator"

ret, _ := calculator.Calculate("1 + 2 + 3", nil)
log.Info(ret) // 6

// 带变量
ret, err := calculator.Calculate("v0*(v1+v2)", map[string]float64{
		"v0": 3,
		"v1": 5,
		"v2": 8,
		"v3": 11,
})
log.Info(ret) // 39
```

### 比较器
```go
import calculator "github.com/smiecj/go_calculator"

ret, _ := calculator.Compare("2 >= 1") // true
```