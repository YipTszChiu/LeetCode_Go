package Search

//53.Ⅱ 0~n-1中缺失的数字
//https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/

func missingNumber(nums []int) int {
	left := 0
	right := len(nums)-1
	var mid int
	for left <= right {
		//先right - left 可以避免right + left上溢
		mid = (left + right) / 2
		if nums[mid] != mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}
