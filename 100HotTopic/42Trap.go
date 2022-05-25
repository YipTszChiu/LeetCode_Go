package _00HotTopic

// 42. 接雨水
// https://leetcode.cn/problems/trapping-rain-water/

// 解法一：暴力破解
func trap(height []int) int {
	// ans 用于存储返回结果
	ans := 0
	length := len(height)
	// 由于左右两个端点不可能储水，因此从 1 遍历到 length-1
	for i := 1; i < length-1; i++ {
		// 每次初始化容器左右两边
		maxLeft, maxRight := 0, 0
		// 遍历容器左端
		for j := i; j >= 0; j-- {
			maxLeft = max(maxLeft, height[j])
		}
		// 遍历容器右端
		for j := i; j < length-1; j++ {
			maxRight = max(maxRight, height[j])
		}
		// 减去本身的高度即为 height[i] 这一格可以储水的量
		ans += min(maxLeft, maxRight) - height[i]
	}

	return ans
}

// 解法二：动态编程
func trap2(height []int) int {
	// ans 用于返回结果
	ans := 0
	length := len(height)
	// 使用两个数组维护关于 height[i] 的左右最大高度
	leftMax, rightMax := make([]int, length), make([]int, length)

	// 初始化 leftMax[0], rightMax[0]
	leftMax[0], rightMax[length-1] = height[0], height[length-1]
	// 从左往右遍历，存储 height[i] 的左端最大高度
	for i := 1; i < length; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	// 从右往左遍历，存储 height[i] 的右端最大高度
	for i := length - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	// 遍历整个 height 数组获取储水容量
	for i := 1; i < length-1; i++ {
		ans += min(leftMax[i], rightMax[i]) - height[i]
	}

	return ans
}

// 解法三：双指针法
func trap3(height []int) int {
	// 初始化双指针
	left, right := 0, len(height)-1
	// 初始化返回值ans
	ans := 0
	// 初始化左右端最大高度
	leftMax, rightMax := 0, 0

	// 开始遍历
	for left < right {
		// 左指针小于右指针，容量由左决定
		if height[left] < height[right] {
			// 如果左指针大于等于左端最大值，更新
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				// 左指针小于左端最大值，容量取决于左指针
				ans += leftMax - height[left]
			}
			left++
		} else { // 左指针大于等于右指针，容量由右决定
			// 右指针大于等于右端最大值，更新
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				// 右指针小于右端最大值，容量取决于右指针
				ans += rightMax - height[right]
			}
			right--
		}
	}

	return ans
}
