package _00HotTopic

import sort2 "sort"

// 406. 根据身高重建队列
// https://leetcode.cn/problems/queue-reconstruction-by-height/?favorite=2cktkvj

func reconstructQueue(people [][]int) [][]int {
	// 将数组进行排序：首先按照身高排列由低到高，当身高相同，则按照前面人数由高到低
	sort2.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})
	// ans 存储返回结果
	ans := make([][]int, len(people))

	// 遍历person数组
	for _, person := range people {
		// person[1]表示在该person前面大于等于person[0]的有person[1]个，因此该person应该排在person[1]+1位
		spaces := person[1] + 1
		// 遍历返回数组
		for i := range ans {
			// 返回数组第 i 位为空，表示可以插入
			if ans[i] == nil {
				spaces--
				// spaces减到0表示该person应该放到该空位置
				if spaces == 0 {
					ans[i] = person
					break
				}
			}
		}
	}
	return ans
}
