package Search

//剑指 Offer 11. 旋转数组的最小数字
//https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/

//基于二分法查找
func minArray(numbers []int) int {
	left := 0
	right := len(numbers)-1

	for left < right {
		mid := left + (right - left) / 2
		//断点左右两边分别递增，如果mid < right，说明从mid~right递增，断点不在mid的右边
		if numbers[mid] < numbers[right] {
			//不可以right = mid-1，因为不能排除mid可能是断点
			right = mid
		} else if numbers[mid] > numbers[right] {
			//mid > right，说明从mid~right不递增，断点在mid+1 ~ right中
			left = mid + 1
		} else {
			//当mid = right 数组中有可能出现连续重复元素，不能排除
			right--
		}
	}

	return numbers[left]
}