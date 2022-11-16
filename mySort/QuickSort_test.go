package mySort

import "testing"

func TestQuickSort(t *testing.T) {
	arr := []int{5, 9, 8, 1, 3, 8, 7, 4, 6, 88}
	res := []int{1, 3, 4, 5, 6, 7, 8, 8, 9, 88}
	QuickSort(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] != res[i] {
			t.Errorf("sort error: %s", arr)
		}
	}
}
