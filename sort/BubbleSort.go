package sort

// 时间复杂度：O(n) - O(n^2) 平均 O(n^2)
// 空间复杂度：O(1)
// 稳定性：√
func BubbleSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
