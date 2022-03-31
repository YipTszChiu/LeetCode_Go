package DivideAndConquer

//剑指 Offer 16. 数值的整数次方
//https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/

func myPow(x float64, n int) float64 {
	//特值处理
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	//负值处理
	if n < 0 {
		x = 1 / x
		n = -n
	}

	temp := myPow(x, n/2)

	//如果是奇数次方，需要再乘一个x，如5次方的temp划分为2次 * 2次，会少一次
	if n%2 == 1 {
		return temp * temp * x
	} else {
		return temp * temp
	}
}
