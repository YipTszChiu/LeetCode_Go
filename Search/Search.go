package Search

import "sort"

//剑指 Offer 53 - I. 在排序数组中查找数字 I
//https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/

func search(nums []int, target int) int {
	//使用go自带的二分查找：注意只能在有序前提下
	//找到target的位置
	left := sort.SearchInts(nums, target)
	if left == len(nums) {
		return 0
	}
	//找到比target大的位置
	right := sort.SearchInts(nums, target+1)
	return right - left
}
