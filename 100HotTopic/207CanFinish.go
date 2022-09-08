package _00HotTopic

// 207. 课程表
// https://leetcode.cn/problems/course-schedule/?favorite=2cktkvj

func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses)
		result  []int
		valid   = true
		dfs     func(u int)
	)
	// 深度优先遍历
	dfs = func(u int) {
		// visited 位置为 1 表示正在遍历
		visited[u] = 1
		// 对 u 的所有边进行遍历
		for _, v := range edges[u] {
			// 所指向的节点未遍历
			if visited[v] == 0 {
				// 对该节点进行深度遍历
				dfs(v)
				// 如果发现已经不符合要求，直接返回
				if !valid {
					return
				}
			} else if visited[v] == 1 { // 所指向的节点正在被遍历，说明有环，不符合要求
				valid = false
				return
			}
		}
		// 结束将该节点置为2表示遍历完成
		visited[u] = 2
		// 将该节点入栈
		result = append(result, u)
	}
	// 遍历所有前置课程初始化图
	for _, info := range prerequisites {
		// 构建图：想要上0先要上1，因此1指向0
		edges[info[1]] = append(edges[info[1]], info[0])
	}
	// 遍历所有未遍历的节点
	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	return valid
}
