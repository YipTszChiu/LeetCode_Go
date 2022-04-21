package DoublePointer

//剑指 Offer 57. 和为s的两个数字
//https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof/

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}

	//新建两个指针
	ptr1, ptr2 := 0, len(nums)-1

	//根据递增序列的特性进行遍历
	for ptr1 < ptr2 {
		//找到和为target直接返回
		if nums[ptr1]+nums[ptr2] == target {
			return []int{nums[ptr1], nums[ptr2]}
		} else if nums[ptr1]+nums[ptr2] < target {
			//两者和小于target，ptr1后移
			ptr1++
		} else {
			//两者和大于target，ptr2前移
			ptr2--
		}
	}
	//没有找到合适的结果
	return nil
}
