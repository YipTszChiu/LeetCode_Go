package binary_search

// 69. x 的平方根
// https://leetcode.cn/problems/sqrtx/description/

// 二分查找法
func mySqrt(x int) int {
	l, r := 0, x

	ans := -1

	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return ans
}
