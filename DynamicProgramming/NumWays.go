package DynamicProgramming

//青蛙跳台阶问题

func numWays(n int) int {
	//递推公式：f(n) = f(n-1) + f(n-2)
	//边界条件：f(0) = 1; f(1) = 1; f(2) = 2
	a, b, c := 1, 1, 0

	for i := 0; i < n; i++ {
		c = (a + b) % (1e9 + 7)
		a = b
		b = c
	}

	return a
}
