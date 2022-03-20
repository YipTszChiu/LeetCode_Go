package DynamicProgramming

//剑指 Offer 49. 丑数
//https://leetcode-cn.com/problems/chou-shu-lcof/

func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	//第一个丑数是1
	dp[1] = 1
	//三个指针，分别代表乘2，乘3，乘5，初始化指向1
	p2, p3, p5 := 1, 1, 1
	//已知第一个丑数是1，因此从第二个开始遍历
	for i := 2; i <= n; i++ {
		//丑数的因数只能有2，3，5，所以寻找新的丑数只能从原本的丑数中乘以2、3、5
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		//题目要求找到第n个，需要对获得的三个丑数排序，取最小的一个为dp[i]
		dp[i] = min(x2, min(x3, x5))
		//取得的x代表了指针p指向的这个数字已经被2、3或5乘过，新的丑数不会再出现同样的数
		//例子：开始指针均指向index = 1，分别乘以2、3、5后最小的数是2，dp[2] = 2，此时p2需要后移指向index = 2，否则还会出现1 * 2 = 2
		if dp[i] == x2 {
			p2++
		}
		if dp[i] == x3 {
			p3++
		}
		if dp[i] == x5 {
			p5++
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
