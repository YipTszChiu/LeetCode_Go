package BitOperation

//剑指 Offer 15. 二进制中1的个数
//https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/

//如6-110, 6 & (6-1) = 110 & 101 = 100HotTopic,每次与n-1会将最后一个1反转为0
func hammingWeight(num uint32) int {
	count := 0
	for ; num > 0; num &= num - 1 {
		count++
	}
	return count
}
