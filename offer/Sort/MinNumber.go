package Sort

import (
	"bytes"
	"sort"
	"strconv"
)

//剑指 Offer 45. 把数组排成最小的数
//https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/

func minNumber(nums []int) string {
	//空值处理
	if nums == nil || len(nums) == 0 {
		return ""
	}

	//将nums转换为string数组
	strs := make([]string, len(nums))
	for i, _ := range nums {
		strs[i] = strconv.Itoa(nums[i])
	}
	//自定义排序规则
	sort.Slice(strs, func(i, j int) bool {
		s12 := strs[i] + strs[j]
		s21 := strs[j] + strs[i]
		return s12 < s21
	})
	//拼接字符串
	var buffer bytes.Buffer
	for i, _ := range strs {
		buffer.WriteString(strs[i])
	}

	return buffer.String()
}
