package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T)  {
	array := []int{39, 2, 5, 23, 54, 12, 78, 34, 45, 40}
	QuickSort(array)
	fmt.Println("快速排序",array)
}
