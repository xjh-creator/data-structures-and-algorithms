package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T)  {
	array := []int{39, 2, 5, 23, 54, 12, 78, 34, 45, 40}
	QuickSort(array)
	for i, v := range array {
		fmt.Printf("下标为%d的值为%d", i, v)
		fmt.Println()
	}
}
