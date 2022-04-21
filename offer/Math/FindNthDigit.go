package Math

import (
	"strconv"
)

func findNthDigit(n int) int {
	// digit 为位数，比如12为两位。start 为数字所述范围的开始数，如10 ~ 99的start为10；count 为所占位数量，如10~99 的位数量为180=2*90
	digit, start, count := 1, 1, 9

	// 位数n > 数字所占位数量count
	// 例子：如 n = 15， n > 9 （9为1~9所占用的9位）
	for n > count {
		// 减去前面已经计算过的位数
		n -= count
		// n 属于下一个范围，因此digit++
		digit++
		// start 为每个范围起始数字，如1，10，100HotTopic
		start *= 10
		// count 推算公式为 digit * start * 9，如 10~99的count = 2 * 10 * 9 = 180
		count = digit * start * 9
	}
	// 得到n所属范围后，计算是哪个数字
	// 例子：n = 15（指向的是12的1），经过循环后，n = 6，指向的数字为 10 + (6-1) / 2 = 12
	num := start + (n-1)/digit
	// 再计算指向的是num中的哪一位
	index := (n - 1) % digit
	// num转换为字符串
	numStr := strconv.Itoa(num)
	// 得到对应字符后转换回int
	return int(numStr[index] - '0')
}
