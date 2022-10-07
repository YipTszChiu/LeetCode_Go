package _00HotTopic

// 448. 找到所有数组中消失的数字
// https://leetcode.cn/problems/find-all-numbers-disappeared-in-an-array/?favorite=2cktkvj

// 哈希表
// 时间复杂度：O(n) 遍历两次长度为n的数组
// 空间复杂度：O(n) 额外使用一个与nums等长的数组
func findDisappearedNumbers(nums []int) (res []int) {
	// 使用一个长度为n的数组记录存在的数字
	hashMap := make([]int, len(nums))
	// 遍历nums记录数字出现次数
	for i := 0; i < len(nums); i++ {
		hashMap[nums[i]-1]++
	}
	// 遍历哈希表返回消失的数字
	for i := 0; i < len(hashMap); i++ {
		if hashMap[i] == 0 {
			res = append(res, i+1)
		}
	}
	return res
}

// 哈希表的优化
// 时间复杂度：O(n) 遍历两次长度为n的数组
// 空间复杂度：O(1) 除了返回数组之外没有额外使用空间
func findDisappearedNumbers2(nums []int) (res []int) {
	// 思路：遍历到一个数字x，将对应的nums[x-1]存储的数值加n；
	// 由于nums中的数字范围是1~n，因此遍历后如果没有超过n的位置说明i+1这个数字缺失；
	// 注意由于遍历某个数字时可能已经被加n，因此需要取模获得原来的值
	n := len(nums)
	// 第一次遍历对相应的数字加n
	for i := 0; i < n; i++ {
		nums[(nums[i]-1)%n] += n
	}
	// 第二次遍历找出缺失的数字
	for i := 0; i < n; i++ {
		if nums[i] <= n {
			res = append(res, i+1)
		}
	}

	return
}
