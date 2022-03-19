package DynamicProgramming

//面试题19. 正则表达式匹配
//https://leetcode-cn.com/problems/zheng-ze-biao-da-shi-pi-pei-lcof/

//考虑正则串末尾：1、普通字符；2、'.'；3、'*'
func isMatch(s string, p string) bool {
	//m，n分别为主串和正则串的长度
	m, n := len(s), len(p)

	//i, j传的是从1开始计数的位置
	//matches函数用于判断主串第i位和正则串第j位能否匹配
	matches := func(i, j int) bool {
		if i == 0 {
			//传入参数从1计数，不存在0位
			return false
		}
		//j-1为字符串中从0开始的位置
		//如果正则串末尾是'.' 可以匹配任意字符
		if p[j-1] == '.' {
			return true
		}
		//如果正则串末尾是普通字符，判断是否与主串匹配
		return s[i-1] == p[j-1]
	}

	//构建m+1行的二位表，多一行是空串
	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		//完成 [m+1, n+1]二维表的构建，多一列是空正则串
		f[i] = make([]bool, n+1)
	}
	//f[0][0]表示空字符串与空正则串能匹配
	f[0][0] = true

	//注意是 <= 因为二维表是 [m+1,n+1]
	//tips：关于f[i][j]，matches[i][j]，i和j表示的是从1计数的位置；关于s[i], p[j]，i和j表示的是从0计数的位置
	for i := 0; i <= m; i++ {
		//j = 0 表示空正则串，与任何非空主串都不匹配
		for j := 1; j <= n; j++ {

			//如果正则串末尾是'*'
			if p[j-1] == '*' {
				// '*'和前面一个字符c组成 'c*'，又因为'*'可以代表0个c，考虑直接无视'c*'看f[i][j-2]
				//tips： 需要 || 的原因是避免已经判断为true的项被下面一步更改
				f[i][j] = f[i][j] || f[i][j-2]
				// 这里认为'c*'至少代表一个c，则需要判断这里的c字符，是否与主串中的字符匹配
				if matches(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}

			} else if matches(i, j) {
				//如果正则串末尾不是'*'，使用matches函数，可以判断是普通字符或者'.'的情况
				f[i][j] = f[i][j] || f[i-1][j-1]
			}

			//最后的else属于末尾既不是'*'，也不匹配，默认值false无需更改
		}
	}

	return f[m][n]
}
