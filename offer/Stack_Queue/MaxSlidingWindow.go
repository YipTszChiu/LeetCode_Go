package Stack_Queue

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 { return nil }
	ans := []int{}
	que := []int{}

	push := func(i int) {
		//que非空，判断单调队列末尾元素是否小于nums[i]，若果是，出队
		/*注意：此处nums[i] 不能 >= que[len(que)-1]，因为如果遇见同样大的元素，
		  如[5,7,6,7]，会将第一个7出队，导致错误
		*/
		for len(que) > 0 && nums[i] > que[len(que)-1] {
			que = que[:len(que)-1]
		}
		//将nums[i]加入到队尾
		que = append(que, nums[i])
	}
	//初始化单调队列
	for i := 0; i < k; i++ {
		push(i)
	}
	//此时已经形成窗口，将第一个窗口的最大值加入ans
	ans = append(ans, que[0])
	//开始遍历剩下元素
	for i := k; i < len(nums); i++ {
		push(i)
		//滑动窗口左边元素过期，如果此时恰好为窗口中最大元素，需要在que中出队
		if que[0] == nums[i-k] {
			que = que[1:]
		}
		ans = append(ans, que[0])
	}
	return ans
}
