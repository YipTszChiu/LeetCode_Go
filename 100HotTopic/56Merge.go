package _00HotTopic

import "sort"

func merge56(intervals [][]int) [][]int {
	// 将 intervals 排序，这样可以保证里面的区间左端都是升序的
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// res 用于返回结果
	res := [][]int{}
	// prev 记录前面完成合并的区间，初始化为 interval[0]
	prev := intervals[0]

	// 遍历整个数组
	for i := 1; i < len(intervals); i++ {
		// 记录当前遍历的区间
		cur := intervals[i]
		// 判断已经合并的区间右端与当前区间左端的大小
		if prev[1] < cur[0] {
			// 不重合
			res = append(res, prev)
			prev = cur
		} else {
			// 有重合
			prev[1] = max(prev[1], cur[1])
		}
	}
	res = append(res, prev)
	return res
}
