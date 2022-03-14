package SearchAndBackchecking

//剑指 Offer 13. 机器人的运动范围
//https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/

//visit二维数组用于维护dfs避免重复遍历
var visit [][]bool

func movingCount(m int, n int, k int) int {
	//初始化visit
	visit = make([][]bool, m)
	for i := 0; i < m; i++ {
		visit[i] = make([]bool, n)
	}
	return dfs2(0, 0, m, n, k)
}

//i,j为当前遍历位， m,n为行列的上限， k为和的限制
func dfs2(i, j, m, n, k int) int {
	//越界，坐标和大于k，已经遍历过 -> 返回0
	if i >= m || j >= n || transferSum(i)+transferSum(j) > k || visit[i][j] {
		return 0
	}
	visit[i][j] = true
	return 1 + dfs2(i+1, j, m, n, k) + dfs2(i, j+1, m, n, k)
}

//返回坐标的和
func transferSum(a int) int {
	sum := 0
	for a > 0 {
		sum += a % 10
		a /= 10
	}
	return sum
}
