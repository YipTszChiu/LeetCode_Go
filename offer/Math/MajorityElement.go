package Math

import "sort"

//剑指 Offer 39. 数组中出现次数超过一半的数字
//https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/

// 方法一：哈希表
func majorityElement(nums []int) int {
	// 创建一个用于查找一个整数出现次数的哈希表
	elemMap := make(map[int]int)
	// 遍历nums每一个元素记录次数
	for _, elem := range nums {
		elemMap[elem]++
		// 当某个元素出现次数超过一半数字时直接跳出循环
		if elemMap[elem] > len(nums)/2 {
			return elem
		}
	}

	return 0
}

// 方法二：排序
func majorityElement2(nums []int) int {
	// 将数组排序，出现次数大于一半的数必然出现在中间
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// 方法三：摩尔投票法
func majorityElement3(nums []int) int {
	// 初始化投票数
	vote := 0
	// 候选人x
	var x int
	// 遍历整个数组
	for _, elem := range nums {
		// 当前票数为0，将最新的元素当作候选人
		if vote == 0 {
			x = elem
		}
		if elem == x {
			vote++
		} else {
			vote--
		}
	}
	// 验证x是否为超过一半的数
	count := 0
	for _, elem := range nums {
		if x == elem {
			count++
		}
	}
	if count > len(nums)/2 {
		return x
	}

	return 0
}
