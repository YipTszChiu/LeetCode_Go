package DivideAndConquer

import (
	"math"
	"strconv"
)

//剑指 Offer 17. 打印从1到最大的n位数
//https://leetcode-cn.com/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/

//需要考虑大数问题
func printNumbers(n int) []int {
	//对0~9做全排列，但是注意首位不应为0，如001，010是不合法的
	//res用于返回结果，长度为10^n-1，因为n位数总共有10^n-1个
	res := make([]int, int(math.Pow10(n)-1))
	//count用于记录已经遍历完成的个数
	count := 0
	//声明并定义dfs函数：index为当前遍历的位置，num存储已经固定的位，digit为总共需要全排列的位数
	var dfs func(index int, num []byte, digit int)

	//首先考虑处理首位，只能为1~9，步骤 ** 处

	//dfs按位遍历实现全排列
	dfs = func(index int, num []byte, digit int) {
		//index为当前遍历位，如果已经到digit总位数，将该排列加入到结果中并结束返回
		if index == digit {
			//将byte数组的字符转为string格式，再转为int格式，这一步是因为LeetCode该题需要返回[]int，正常来说大数返回[]string
			temp, _ := strconv.Atoi(string(num))
			res[count] = temp
			count++
			return
		}
		//固定下一位进行遍历，注意此时已经从第二位开始，可以为'0'
		for i := '0'; i <= '9'; i++ {
			//num = append(num, byte(i))    记录一次错误的写法：由于使用的是[]byte数组，是无法append的
			num[index] = byte(i)
			dfs(index+1, num, digit)
		}
	}

	//** 首位处理并进行全排列
	//digit为数字的总位数，从1位遍历到需要输出的n位
	for digit := 1; digit <= n; digit++ {
		//**首位处理
		for first := '1'; first <= '9'; first++ {
			//初始化一个[]byte数组，长度为digit位
			num := make([]byte, digit)
			//将first转为byte形式固定到第一位
			num[0] = byte(first)
			//调用dfs对后面的位进行遍历
			dfs(1, num, digit)
		}
	}

	return res
}
