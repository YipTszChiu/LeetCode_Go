package DynamicProgramming

//斐波那契数列

func fib(n int) int {
	//递推公式：f(n) = f(n-1) + f(n-2)
	//边界条件：f(0) = 0, f(1) = 1
	if n < 2 {
		return n
	}

	a, b, c := 0, 0, 1
	for i := 2; i <= n; i++ {
		a = b
		b = c
		c = (a + b) % (1e9 + 7)
	}
	return c
}
