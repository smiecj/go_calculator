package test

import (
	"log"
	"testing"

	"github.com/smiecj/go_calculator/service"
	"github.com/stretchr/testify/require"
)

// 测试计算纯数字
func TestPureNumber(t *testing.T) {
	ret, err := service.Calculate("1 + 2 + 3", map[string]float64{})
	if nil != err {
		log.Printf("计算失败，错误原因: %s", err.Error())
	} else {
		log.Printf("计算成功，结果: %.2f", ret)
	}
	require.Equal(t, nil, err)
	require.Equal(t, float64(6), ret)

	// 稍微复杂一点的计算
	ret, err = service.Calculate("6 * ( 2 + 3)", map[string]float64{})
	require.Equal(t, nil, err)
	require.Equal(t, float64(30), ret)

	ret, err = service.Calculate("(4490-0)*100/4490", map[string]float64{})
	require.Equal(t, nil, err)
	require.Equal(t, float64(100), ret)
}

// 测试变量计算
func TestVariable(t *testing.T) {
	variableMap := map[string]float64{
		"v0": 3,
		"v1": 5,
		"v2": 8,
		"v3": 11,
	}
	emptyVariableMap := map[string]float64{}

	ret, err := service.Calculate("v0 + v1", variableMap)
	if nil != err {
		log.Printf("计算失败，错误原因: %s", err.Error())
	} else {
		log.Printf("计算成功，结果: %.2f", ret)
	}
	require.Equal(t, nil, err)
	require.Equal(t, float64(8), ret)

	// 稍微复杂一点的计算
	ret, err = service.Calculate("v0 * (  v1 + v2 )", variableMap)
	require.Equal(t, nil, err)
	require.Equal(t, float64(39), ret)

	ret, err = service.Calculate("v1 * (5 + v3)", variableMap)
	require.Equal(t, nil, err)
	require.Equal(t, float64(80), ret)

	ret, err = service.Calculate("(v0+v1)/(v2+v3)", variableMap)
	require.Equal(t, nil, err)
	log.Printf("[test] calculate ret: %.2f", ret)

	ret, err = service.Calculate("4", emptyVariableMap)
	require.Equal(t, nil, err)
	log.Printf("[test] calculate ret: %.2f", ret)
}
