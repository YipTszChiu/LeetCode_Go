package BitOperation

//剑指 Offer 56 - II. 数组中数字出现的次数 II
//https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/

//需要自己画图按位推理
func singleNumber(nums []int) int {
	two, one := 0, 0
	for _, elem := range nums {
		one = (one ^ elem) & (^two)
		two = (two ^ elem) & (^one)
	}
	return one
}
