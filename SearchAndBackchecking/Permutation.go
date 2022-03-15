package SearchAndBackchecking

import (
	"sort"
)

//剑指 Offer 38. 字符串的排列
//https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/solution/zi-fu-chuan-de-pai-lie-by-leetcode-solut-hhvs/

//方法一：回溯
func permutation(s string) (ans []string) {
	//将string转换为byte数组
	t := []byte(s)
	//将t顺序排列，即相同的字符会放在相邻位置
	sort.Slice(t, func(i, j int) bool {
		return t[i] < t[j]
	})

	//创建一个长度为length的[]byte
	length := len(t)
	perm := make([]byte, 0, length)
	//创建一个长度为length的[]bool且初始化为false
	vis := make([]bool, length)

	//声明并定义backtrack
	var backtrack func(int)
	//i表示正在填充第i位
	backtrack = func(i int) {
		//当i == length,到达递归终点
		if i == length {
			//perm数组转换为string字符串追加到结果
			ans = append(ans, string(perm))
			return
		}
		for j, b := range vis {
			//跳过该次循环条件：
			//1. b == true 即 该位字符的vis == true，已经使用过
			//2. t[j-1] == t[j] 即 前后字符相同，且前面vis[j-1] == true已经使用过
			if b || j > 0 && vis[j-1] && t[j-1] == t[j] {
				continue
			}
			//否则将vis[j]置为true,表示已经用过第j位,并将t[j]加入到perm中
			vis[j] = true
			perm = append(perm, t[j])
			//递归调用
			backtrack(i + 1)
			//回溯，将末尾一位去除并重置vis[j]为false
			perm = perm[:len(perm)-1]
			vis[j] = false
		}
	}

	//从第0位开始递归
	backtrack(0)

	return
}
