package _00HotTopic

// 17. 电话号码的字母组合
// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

// 构建一个 数字 - 字符数组 的映射，如 2 - [a,b,c]
var hashMap map[byte][]byte = map[byte][]byte{
	'2': []byte{'a', 'b', 'c'},
	'3': []byte{'d', 'e', 'f'},
	'4': []byte{'g', 'h', 'i'},
	'5': []byte{'j', 'k', 'l'},
	'6': []byte{'m', 'n', 'o'},
	'7': []byte{'p', 'q', 'r', 's'},
	'8': []byte{'t', 'u', 'v'},
	'9': []byte{'w', 'x', 'y', 'z'},
}

// res用于返回结果
var res []string

func letterCombinations(digits string) []string {
	// 空值处理
	if len(digits) == 0 {
		return nil
	}
	res = []string{}
	backtrack(digits, 0, "")
	return res
}

func backtrack(digits string, index int, combination string) {
	// 遍历结束，将结果加入返回数组中
	if index == len(digits) {
		res = append(res, combination)
	} else {
		digit := digits[index]
		letters := hashMap[digit]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			backtrack(digits, index+1, combination+string(letters[i]))
		}
	}

}
