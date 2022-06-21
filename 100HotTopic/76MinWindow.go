package _00HotTopic

import "math"

func minWindow(s string, t string) string {
	// 使用两个哈希表判断存在的字符
	ori, cnt := map[byte]int{}, map[byte]int{}
	// 遍历字符串 t 记录需要包含的字符
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	len := math.MaxInt32
	// 初始化滑动窗口的左右指针
	ptrL, ptrR := -1, -1

	// check 函数用于判断 s 中是否包含足够数量的 t 中的字符
	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}

	// 遍历字符串 s
	for l, r := 0, 0; r < sLen; r++ {
		// 判断 s 中一个字符在 t 中是否含有，有的话 cnt 中该字符的映射+1
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		// 调用 check 检查 s 是否包含足够数量的 t 中字符,且左指针为超过右指针（注意这里是一边 r 自增，一边check，所以要判断 l 和 r）
		for check() && l <= r {
			if r-l+1 < len {
				// s 中包含 t 中足够数量的字符， 且比当前窗口短，则更新窗口
				len = r - l + 1
				// 调整滑动窗口指针变为当前的长度
				ptrL, ptrR = l, l+len
			}
			// 检查 t 中是否含有 s[l] 这一个字符
			// 如果有则 cnt 中对应字符映射-1，这样会导致下次循环 check 返回 false，由于最左端含有需要的字符，此时窗口为最小
			if _, ok := ori[s[l]]; ok {
				// 将记录 s 中该字符的映射-1
				cnt[s[l]] -= 1
			}
			// 如果没有，不需要-1，check仍然通过，这样可以缩小窗口
			// 左指针右移，缩短距离
			l++
		}
	}

	if ptrL == -1 {
		return ""
	}

	return s[ptrL:ptrR]
}
