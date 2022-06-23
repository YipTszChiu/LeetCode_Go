package _00HotTopic

// 78. 子集
// https://leetcode.cn/problems/subsets/

// 方法一：掩码
func subsets1(nums []int) (ans [][]int) {
	n := len(nums)
	// nums 有 n 位，每一位通过二进制决定是否在子集中出现，如 001 则第三位出现
	for mask := 0; mask < 1<<n; mask++ {
		set := []int{}
		// 遍历 nums
		for i, v := range nums {
			// nums 一共有多少个数，mask就一共需要右移几次（包括0次），右移后如果该位置的掩码为 1，将其加入到 set
			if mask>>i&1 > 0 {
				set = append(set, v)
			}
		}
		// 一次掩码完整遍历结束，将 set 加入到 ans
		ans = append(ans, set)
	}

	return
}

// 方法二：递归
func subsets2(nums []int) (ans [][]int) {
	set := []int{}
	// 声明并定义递归函数
	var dfs func(int)
	dfs = func(cur int) {
		// 如果已经递归完整个 nums
		if cur == len(nums) {
			// 将当前的 set 加入到 ans
			ans = append(ans, append([]int{}, set...))
			return
		}
		// 考虑选择当前 cur 位
		set = append(set, nums[cur])
		// 继续递归下一位
		dfs(cur + 1)
		// 回溯时删除 cur 位
		set = set[:len(set)-1]
		// 考虑不选择当前 cur 位
		dfs(cur + 1)
	}
	dfs(0)
	return
}
