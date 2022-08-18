package _00HotTopic

// 136. 只出现一次的数字
// https://leetcode.cn/problems/single-number/

// 哈希表法
// 时间复杂度：遍历 nums 构建 numMap 时间 O(n)，再次遍历同样是 O(n)，时间复杂度为 O(n)
// 空间复杂度：使用了额外空间 map，空间复杂度为 O(n)
func singleNumber(nums []int) int {
	// 使用 map 记录一个数字是否出现过
	numMap := map[int]int{}
	// 遍历 nums 并修改 numMap
	for _, num := range nums {
		numMap[num]++
	}
	// 再次遍历 nums 找到只出现一次的数字
	for _, num := range nums {
		if numMap[num] == 1 {
			return num
		}
	}
	return 0
}

// 位运算异或法：题目明确说明其他数字均出现两次，只有一个数字出现一次，并要求 线性时间复杂度与 不使用额外空间，因此可以使用异或运算
func singleNumber2(nums []int) int {
	// 任何数异或 0 还是本身
	res := 0
	// 遍历 nums 与 res 异或
	for _, num := range nums {
		res ^= num
	}
	return res
}
