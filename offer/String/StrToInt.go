package String

import (
	"math"
	"strings"
)

//复习：1
func strToInt(str string) int {
	//去除多余空格
	str = strings.TrimSpace(str)
	result := 0
	sign := 1

	for i, v := range str {
		//表示数字的字符
		if v >= '0' && v <= '9' {
			//已有结果左移一位（十进制），加上当前位：ASCII码的整数值 = v - '0'
			result = result*10 + int(v-'0')
		} else if v == '+' && i == 0 { //关键：必须加上 i == 0 的条件，只有第一位是符号才合法
			//处理正负号
			sign = 1
		} else if v == '-' && i == 0 {
			sign = -1
		} else {
			break
		}
	}
	//根据规则，如果溢出则返回MaxInt或MinInt
	if result > math.MaxInt32 {
		if sign == -1 {
			return math.MinInt32
		}
		return math.MaxInt32
	}
	//不要遗忘符号位
	return sign * result
}
