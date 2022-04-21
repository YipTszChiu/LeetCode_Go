package _00HotTopic

// 1. 两数之和
// https://leetcode-cn.com/problems/two-sum/

// 哈希表方法
func twoSum(nums []int, target int) []int {
	// 创建一个哈希表保存遍历过的内容
	hashTable := make(map[int]int)
	// 遍历整个数组
	for i, x := range nums {
		// 查看map中是否有target-x的映射，如果有，说明符合两数之和条件，返回数组
		if elem, ok := hashTable[target-x]; ok {
			return []int{i, elem}
		}
		// 如果没有，存入nums数组所存值，与其index的映射
		hashTable[x] = i
	}
	return nil
}
