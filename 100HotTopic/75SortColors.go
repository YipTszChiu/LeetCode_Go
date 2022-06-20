package _00HotTopic

// 75. 颜色分类
// https://leetcode.cn/problems/sort-colors/

// 双指针法
func sortColors(nums []int) {
	// 双指针初始指向数组头部
	p0, p1 := 0, 0
	// 遍历数组
	for i, c := range nums {
		if c == 0 {
			// 交换这个 “0” 与指针指向的位置
			nums[i], nums[p0] = nums[p0], nums[i]
			// p0与p1不指向同一个位置，说明已经存有 1 在前面，当发生 0 与 num[i]交换时，会导致本身在 0 后的 1 被放到nums[i]
			// 因此需要将这个 1 放回来
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p0++
			p1++
		} else if c == 1 {
			// 对 “1” 进行交换
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
		// 对 2 不需要进行处理，经过上面的交换后会排在所有 0 和 1 之后
	}
}
