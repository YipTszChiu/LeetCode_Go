package _00HotTopic

// 11. 盛最多水的容器
// https://leetcode-cn.com/problems/container-with-most-water/

// 双指针法
func maxArea(height []int) int {
	// 初始化左右指针分别指向数组开头以及末尾
	left, right := 0, len(height)-1
	// 初始化最大容积 = 短板 * 桶高
	maxA := min2(height[left], height[right]) * (right - left)
	// 对数组进行遍历，每次移动较小的木板
	for left < right {
		// 左指针为短板，左指针右移
		if height[left] <= height[right] {
			left++
		} else {
			// 右指针为短板，右指针左移
			right--
		}
		// 计算指针移动后的容积，如果大于最大容积则赋值给maxA
		temp := min2(height[left], height[right]) * (right - left)
		if temp > maxA {
			maxA = temp
		}
	}
	return maxA
}

func min2(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
