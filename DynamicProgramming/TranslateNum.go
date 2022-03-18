package DynamicProgramming

import "strconv"

//剑指 Offer 46. 把数字翻译成字符串
//https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/

func translateNum(num int) int {
	//将int转换为字符串
	str := strconv.Itoa(num)
	p, q, r := 0, 0, 1

	for i := 0; i < len(str); i++ {
		p, q, r = q, r, 0
		//dp[i-1]是无论如何都能翻译成功的，dp[i] += dp[i-1]
		r += q
		//取前缀时i-1有可能为负
		if i == 0 {
			continue
		}
		//取 i-1, i两位组成数字，判断是否能翻译：00，01无法翻译
		pre := str[i-1 : i+1]
		//将pre转换为整数
		preInt, _ := strconv.Atoi(pre)
		//如果在范围内则可以翻译，dp[i] = dp[i-1] + dp[i-2]
		if preInt >= 10 && preInt <= 25 {
			r += p
		}
	}

	return r
}
