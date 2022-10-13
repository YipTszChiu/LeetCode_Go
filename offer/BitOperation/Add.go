package BitOperation

//剑指 Offer 65. 不用加减乘除做加法
//https://leetcode-cn.com/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/

// 当前位相加等于异或 a ^= b，进制位置等于相与
func Add(a int, b int) int {
	for b > 0 {
		c := (a & b) << 1 //c为进位
		a ^= b            //a为当前位，与b异或后存到a
		b = c             //将进制位赋给b，如果b不为0需要继续将进制位加上
	}

	return a
}
