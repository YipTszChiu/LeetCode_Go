package DynamicProgramming

//剑指 Offer 63. 股票的最大利润
//https://leetcode-cn.com/problems/gu-piao-de-zui-da-li-run-lcof/

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	//初始化最小价格和最大利润
	minPrice := prices[0]
	maxPro := 0

	for i := 0; i < len(prices); i++ {
		//更新最小价格
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxPro {
			//更新最大利润
			maxPro = prices[i] - minPrice
		}
	}

	return maxPro
}
