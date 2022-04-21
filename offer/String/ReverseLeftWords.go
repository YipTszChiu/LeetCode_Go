package String

func reverseLeftWords(s string, n int) string {
	b := []byte(s)

	result := b[n:]
	for i := 0; i < n; i++ {
		result = append(result, b[i])
	}

	return string(result)
}

//复习时候发现更快的做法
//func reverseLeftWords(s string, n int) string {
//	res := []byte(s)
//	temp := res[:n]
//	res = res[n:]
//	res = append(res, temp...)
//
//	return string(res)
//}
