package Math

// 剑指 Offer 66. 构建乘积数组
// https://leetcode-cn.com/problems/gou-jian-cheng-ji-shu-zu-lcof/

func constructArr(a []int) []int {
	// 获取数组长度
	length := len(a)
	// 空值处理
	if length == 0 {
		return nil
	}

	// 构造要求的数组
	b := make([]int, length)
	// b[0] = 1 * a[1] * ……，因此初始化b[0]
	b[0] = 1

	// 第一次循环构建下三角的元素
	for i := 1; i < length; i++ {
		b[i] = b[i-1] * a[i-1]
	}
	// 第二次循环构建上三角的元素
	temp := 1 // temp用于暂存乘积，用于构建上三角
	for i := length - 2; i >= 0; i-- {
		// b[n-1] = a[0] * a[1] * …… * a[n-2] * 1，末位乘1，因此无需计算，遍历从length -2 开始
		temp *= a[i+1] // a[length-1]
		b[i] *= temp
	}

	return b
}
