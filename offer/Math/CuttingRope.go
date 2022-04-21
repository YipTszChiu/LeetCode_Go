package Math

import "math"

// 剑指 Offer 14- I. 剪绳子
// https://leetcode-cn.com/problems/jian-sheng-zi-lcof/

// 方法一：数学推导
func cuttingRope(n int) int {
	// 由于规定切后段数 m > 1，因此长为 2 = 1 + 1，长为 3 = 1 + 2，乘积为n-1直接返回
	if n <= 3 {
		return n - 1
	}
	// 推论：所有段数为3时乘积最大
	// 计算余数以及商
	remainder := n % 3
	count := n / 3
	// 正好整除3
	if remainder == 0 {
		return int(math.Pow(float64(3), float64(count)))
	}
	// 余数为1，则将一段3分成2+2，乘积更大 ： 1+3 = 2+2
	if remainder == 1 {
		return int(math.Pow(float64(3), float64(count-1)) * 4)
	}
	// 余数为2则最后一段2单独留下
	return int(math.Pow(float64(3), float64(count)) * 2)
}

// 方法二：动态规划
func cuttingRopeDP(n int) int {
	// 由于规定切后段数 m > 1，因此长为 2 = 1 + 1，长为 3 = 1 + 2，乘积为n-1直接返回
	if n <= 3 {
		return n - 1
	}

	// dp思路：砍成两段的最大乘积是 f(n) = f(x)*f(n-x)
	// 另外绳子长度计算时应该按照长度计算，比如dp[2]的结果应该是 1+1=2， 乘积为1，但计算时候要按dp[2]=2算
	// 初始化一个数组，dp[0]不使用，为了方便对照直接从dp[1]开始
	dp := make([]int, n+1)
	for i := 1; i <= 3; i++ {
		dp[i] = i
	}
	temp := 1
	for i := 4; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			temp = max(temp, dp[j]*dp[i-j])
		}
		// 将最大值赋给dp[i]，并将temp置零
		dp[i] = temp
		temp = 0
	}
	return dp[n]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// 剑指 Offer 14- II. 剪绳子 II
// https://leetcode-cn.com/problems/jian-sheng-zi-ii-lcof/
func cuttingRope3(n int) int {
	// 题目条件 n >= 2
	if n <= 3 {
		return n - 1
	}
	// a 为 n 除以 3 的商，b 为 n 除以 3 的余数
	a, b := n/3, n%3
	// 余数为0，整除3说明可以均分为3的线段
	if b == 0 {
		return pow(3, a) % (1e9 + 7)
	}
	// 余数为1，有一段需要拆分成2+2
	if b == 1 {
		return pow(3, a-1) * 4 % (1e9 + 7)
	}
	// 余数为2，保留剩下这段2
	return pow(3, a) * 2 % (1e9 + 7)
}

// 由于题目要求余1e9+7，说明是比较大的数，需要使用快速幂
func pow(x, n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	temp := pow(x, n/2)
	if n%2 == 1 {
		return temp * temp * x % (1e9 + 7)
	} else {
		return temp * temp % (1e9 + 7)
	}
}
