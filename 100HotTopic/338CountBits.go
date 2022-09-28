package _00HotTopic

// 338. 比特位计数
// https://leetcode.cn/problems/counting-bits/

// Brian Kernighan 算法
// 对于任意整数，令 x & (x-1) 可以将 x 的二进制表示的最后一个 1 变成 0
func oneCount(x int) (ones int) {
	for ; x > 0; x &= x - 1 {
		ones++
	}
	return
}

// 时间复杂度：O(nlogn)
// 空间复杂度：O(1) 只有返回数组
func countBits1(n int) []int {
	bits := make([]int, n+1)
	for i := 0; i <= n; i++ {
		bits[i] = oneCount(i)
	}
	return bits
}

// 动态规划
// 最高有效位：令小于x的最大2的n次幂为y，如 x = 110，则 y = 100；x = 1010， y = 1000
// 如何判断最高有效位：由于最高有效位只有最高位为1，因此使用 x & (x-1) == 0
// dp[i] 表示 i 二进制表示的 1 的个数，则 dp[i] = dp[i-highBit] + 1

// 时间复杂度：O(n)
// 空间复杂度：O(1)
func countBits2(n int) []int {
	bits := make([]int, n+1)
	highBit := 0
	for i := 1; i < len(bits); i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		bits[i] = bits[i-highBit] + 1
	}
	return bits
}
