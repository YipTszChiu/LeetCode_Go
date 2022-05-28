package _00HotTopic

import "sort"

// 49. 字母异位词分组
// https://leetcode.cn/problems/group-anagrams/

func groupAnagrams(strs []string) [][]string {
	// 空值处理
	if len(strs) == 0 {
		return nil
	}

	hashMap := map[string][]string{}
	for _, str := range strs {
		// 将 strs string类型数组中每个元素转为 slice
		s := []byte(str)
		// 使用自带的 sort 包，自定义规则进行从小到大排序
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		// 将排序好的结果转回 string 类型
		sortedStr := string(s)
		// 在对应的映射中增加该 string
		hashMap[sortedStr] = append(hashMap[sortedStr], str)
	}
	// res 用于返回结果
	res := make([][]string, 0, len(hashMap))
	// 将 hashMap 中的结果复制到 res 中
	for _, v := range hashMap {
		res = append(res, v)
	}

	return res
}
