package Search

//剑指 Offer 03. 数组中重复的数字
//https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/

//法一：空间复杂度O(n)
func findRepeatNumber(nums []int) int {
	//arr记录已有数字
	arr := make([]int, len(nums))
	//遍历nums，对应数字的arr[i++]
	for i := 0; i < len(nums); i++ {
		arr[nums[i]]++
	}
	//遍历arr
	for i := 0; i < len(arr); i++ {
		if arr[i] > 1 {
			return i
		}
	}

	return -1
}

//法二： 空间复杂度O(1)——根据数字范围从0 ~ n-1
func findRepeatNumber2(nums []int) int {
	//从头遍历整个数组
	for i := 0; i < len(nums); {
		//当前位已经为i，如 nums[0] == 0
		if nums[i] == i {
			//注意必须在此处i才自增，不能放到循环的第三个参数中
			//原因：交换之后原本这个位置不一定满足nums[i] == i，仍需要继续交换到符合条件才i++
			i++
			continue
		}
		//当前位的值为nums[i] (设x = nums[i])，x 应该被放到数组的第x位置，即 nums[x] == x
		if nums[i] == nums[nums[i]] {
			//如果第x位置已经为x，说明x重复，返回 x 即 nums[i]
			return nums[i]
		}
		//如果第x位不为x，则交换x和当前位置
		temp := nums[i]
		nums[i] = nums[temp]
		nums[temp] = temp
	}

	return -1
}
