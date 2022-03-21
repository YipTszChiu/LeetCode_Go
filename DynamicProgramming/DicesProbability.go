package DynamicProgramming

//剑指 Offer 60. n个骰子的点数
//https://leetcode-cn.com/problems/nge-tou-zi-de-dian-shu-lcof/

//逆向递推
func dicesProbability(n int) []float64 {
	//可能的总和数有 5n+1 项，如1个骰子1~6，2个骰子2~12

	//dp[i][j]代表：总数i个骰子，和为j的概率，则总共有n行, 每行5n+1列  tips：不是一个严格的i行j列
	dp := make([][]float64, n)
	for i := range dp {
		dp[i] = make([]float64, 5*(i+1)+1)
	}

	//初始化第一行
	for i := range dp[0] {
		dp[0][i] = float64(1) / float64(6) //注意必须带float64()，否则直接1/6 = 0
	}

	//正向递推 dp[i][j] = dp[i-1][j-k] * 1/6，k 取 1~6
	//说明：i个骰子能摇出和为j的方案，只有在i-1个骰子摇出j-1 ~ j-6才能出现，而j-1方案中摇出1才能达到和为j，其中的概率为1/6
	//bug：当n较小时，正向递推j-k明显会越界，如 i = 2，j = 10， j-2 = 8，单个骰子无法摇出8

	//逆向递推 dp[i-1][j] 只会影响 dp[i][j+1]~dp[i][j+6] 的概率
	//第0行单个骰子已经初始化，从第1行开始遍历
	for i := 1; i < n; i++ {
		//逆向递推计算第i-1行对第i行的贡献
		for j := range dp[i-1] {
			for k := 0; k < 6; k++ {
				//由于是不严格的二维数组，可以用以下的例子理解：
				//当 i = 1，表示的是 2 个骰子，列数（即和数）是 2~12 ， dp[1][0] 就代表2个骰子摇出2的概率
				dp[i][j+k] += dp[i-1][j] * (float64(1) / float64(6)) //注意必须带float64()，否则直接1/6 = 0
			}
		}
	}

	return dp[n-1]
}
