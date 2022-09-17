package _00HotTopic

import "math"

// 279. 完全平方数
// https://leetcode.cn/problems/perfect-squares/?favorite=2cktkvj

// 动态规划
// 时间复杂度：O(n根号n)
// 空间复杂度：O(n)
func numSquares(n int) int {
	// dp[i] 表示 i 最少需要多少个平方数来表示，dp[0] = 0
	dp := make([]int, n+1)
	// dp[i] = 1 + min(dp[i-j^2])
	for i := 1; i < len(dp); i++ {
		minCount := math.MaxInt
		// 理解为将 i 拆分为 (j*j) + (i-j*j)，那么 dp[i-j*j] 可以获取后半部分最少需要多少个数表示，再加上 (j*j) 本身这个平方数（+1）
		for j := 1; j*j <= i; j++ {
			minCount = min(minCount, dp[i-j*j])
		}
		dp[i] = 1 + minCount
	}
	return dp[n]
}

// 数学四平方和定理
// 判断是否为完全平方数
func isPerfectSquare(x int) bool {
	y := int(math.Sqrt(float64(x)))
	return y*y == x
}

// 判断是否能表示为 4^k*(8m+7)
// 具体看https://leetcode.cn/problems/perfect-squares/solution/wan-quan-ping-fang-shu-by-leetcode-solut-t99c/
func checkAnswer4(x int) bool {
	for x%4 == 0 {
		x /= 4
	}
	return x%8 == 7
}

func numSquares2(n int) int {
	if isPerfectSquare(n) {
		return 1
	}
	if checkAnswer4(n) {
		return 4
	}
	for i := 1; i*i <= n; i++ {
		j := n - i*i
		if isPerfectSquare(j) {
			return 2
		}
	}
	return 3
}
