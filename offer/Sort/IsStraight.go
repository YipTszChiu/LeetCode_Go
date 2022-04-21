package Sort

import "sort"

//剑指 Offer 61. 扑克牌中的顺子
//https://leetcode-cn.com/problems/bu-ke-pai-zhong-de-shun-zi-lcof/

func isStraight(nums []int) bool {
	//空值处理
	if len(nums) == 0 || nums == nil {
		return false
	}

	//将原数组排序
	sort.Ints(nums)

	//遍历数组，记录joker的个数，并且检查是否有重复的数字
	joker := 0
	for i := 0; i < len(nums)-1; i++ {
		//注意此处需要len(nums)-1，否则nums[i+1]会越界；而如果存在0必定在前两位，末尾不会为0，因此循环少一步不影响判断结果
		if nums[i] == 0 {
			joker++
		} else if nums[i] == nums[i+1] {
			return false
		}
	}

	//当有joker时，由于joker可当任意数，只需要考虑剩下的数是否组成顺子
	//用末尾最大的数减去除大小王以外最小的数，如果是顺子会有以下式子成立
	return nums[len(nums)-1]-nums[joker] < 5

	//存在错误解答：当时写的判断条件为 == 5-(joker+1)，而实际上有可能除了大小王以外的数字不连续 如：0,0,1,2,5可以组成顺子，但式子不成立
}
