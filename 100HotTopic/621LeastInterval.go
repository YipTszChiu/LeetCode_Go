package _00HotTopic

import "math"

// 621. 任务调度器
// https://leetcode.cn/problems/task-scheduler/?favorite=2cktkvj

func leastInterval(tasks []byte, n int) (minTime int) {
	// cnt记录任务的次数
	cnt := map[byte]int{}
	for _, t := range tasks {
		cnt[t]++
	}
	// nextValid：该项任务下一次可以执行的时间；rest：该项任务剩余的次数
	nextValid := make([]int, 0, len(cnt))
	rest := make([]int, 0, len(cnt))
	// 初始化：nextValid 均为 1，rest 为前面 cnt 统计到的数量
	for _, c := range cnt {
		nextValid = append(nextValid, 1)
		rest = append(rest, c)
	}

	for range tasks {
		minTime++
		minNextValid := math.MaxInt64
		// 遍历所有剩余任务，选择一个最早的任务
		for i, r := range rest {
			if r > 0 && nextValid[i] < minNextValid {
				// 更改最小合法时间
				minNextValid = nextValid[i]
			}
		}
		// 如果任务最早合法时间大于当前时间，即所有任务待机，需要跳过
		if minNextValid > minTime {
			minTime = minNextValid
		}
		best := -1
		// 遍历所有剩余任务，选择一个任务执行
		for i, r := range rest {
			// 选择剩余次数最多的任务
			if r > 0 && nextValid[i] <= minTime && (best == -1 || r > rest[best]) {
				best = i
			}
		}
		// 执行该任务：该任务的下次可执行时间调整，剩余次数调整
		nextValid[best] = minTime + n + 1
		rest[best]--
	}
	return
}
