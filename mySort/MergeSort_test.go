package mySort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{5, 9, 8, 1, 3, 8, 7, 4, 6, 88}
	res := []int{1, 3, 4, 5, 6, 7, 8, 8, 9, 88}
	ans := MergeSort(arr)
	fmt.Println(ans)
	for i := 0; i < len(ans); i++ {
		if ans[i] != res[i] {
			t.Errorf("sort error: %v", arr)
			break
		}
	}
}
