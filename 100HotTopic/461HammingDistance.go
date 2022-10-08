package _00HotTopic

import "math/bits"

// 461. 汉明距离
// https://leetcode.cn/problems/hamming-distance/

// 内置计数函数
func hammingDistance(x, y int) int {
	return bits.OnesCount(uint(x ^ y))
}

// 通过异或以及Brian Kernighan算法
func hammingDistance1(x int, y int) int {
	// 将 x 与 y 进行异或就可以得到有多少位不同
	x ^= y
	// 循环执行 x &= x-1 判断有多少位不同
	count := 0
	for x > 0 {
		x &= (x - 1)
		count++
	}

	return count
}
