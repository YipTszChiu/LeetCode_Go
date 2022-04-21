package Math

// 剑指 Offer 43. 1～n 整数中 1 出现的次数
// https://leetcode-cn.com/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof/

func countDigitOne(n int) int {
	// 将整数n按位分割，可以分为 XnXn-1Xn-2……X0
	// 当前位 Xi设置为cur，高位Xn……Xi+1为high，低位Xi-1……X0为low，digit为cur的数位，是10^i，如个位是1，十位是10，百位100
	// 初始化各参数，count用于返回计数结果
	count := 0
	cur := n % 10
	high := n / 10
	low := 0
	digit := 1

	// 循环条件：high 和 cur 同时为 0 时跳出
	for high != 0 || cur != 0 {
		if cur == 0 {
			count += high * digit
		} else if cur == 1 {
			count += high*digit + low + 1
		} else {
			count += (high + 1) * digit
		}
		// 当前位往前进一位
		low += cur * digit
		cur = high % 10
		high /= 10
		digit *= 10
	}

	return count
}
