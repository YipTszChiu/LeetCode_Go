package binary_search

// 367. 有效的完全平方数
// https://leetcode.cn/problems/valid-perfect-square/

func isPerfectSquare(num int) bool {
	l, r := 0, num

	for l <= r {
		mid := l + (r-l)/2
		temp := mid * mid
		if temp == num {
			return true
		} else if temp > num {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return false
}
