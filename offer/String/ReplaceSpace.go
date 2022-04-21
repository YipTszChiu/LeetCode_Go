package String

func replaceSpace(s string) string {
	x := []byte(s)
	result := []byte{}

	for i := 0; i < len(x); i++ {
		if x[i] == ' ' {
			//注意“...”的用法之一：传参时候将参数打散
			//如这里append一次只能追加一个byte，使用...将[]byte("%20")打散
			result = append(result, []byte("%20")...)
		} else {
			result = append(result, x[i])
		}
	}

	return string(result)
}
