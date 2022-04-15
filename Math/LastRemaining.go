package Math

// 剑指 Offer 62. 圆圈中最后剩下的数字
// https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/

// 约瑟夫环问题
func lastRemaining(n int, m int) int {
	if n == 1 {
		return 0
	}

	x := 0
	for i := 2; i <= n; i++ {
		x = (x + m) % i
	}
	return x
}
