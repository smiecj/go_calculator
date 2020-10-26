# go_calculator
使用go 写的计算器

使用方式：可参考test/calculate_test.go 中的代码

第一种方式：纯数值计算
如: 3 * (1 + 2)

第二种方式：数值+变量计算
v1 * (2 + v3)

注意第二种方式需要保证每一个变量的值都在valueMap 当中有，否则最后Calculate 方法会返回error