package BitOperation

//剑指 Offer 56 - I. 数组中数字出现的次数
//https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/

func singleNumbers(nums []int) []int {
	//temp作为异或的基础，任何数异或0等于其本身
	temp := 0
	//由于除了x，y两个数以外，其他数均出现了两次，而一个数异或自己等于0
	//整组异或后的结果相当于x ^= y
	for _, elem := range nums {
		temp ^= elem
	}
	//由于x和y不同，至少会出现一位相异，即temp至少有一位为1
	//找到这个1
	div := 1
	for temp&div == 0 {
		div <<= 1
	}

	x, y := 0, 0
	//根据这一位1进行分组，由于这一位是通过x，y异或得到的结果，因此x，y这一位必然不同，会被分到不同组
	//而同组内其他数字由于都出现了两次，两组分别异或的结果相当于x和y本身
	for _, elem := range nums {
		if elem&div == 0 {
			x ^= elem
		} else {
			y ^= elem
		}
	}

	return []int{x, y}
}
