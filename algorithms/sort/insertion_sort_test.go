package sort

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T)  {
	array := []int{1,4,5,8,9,2,10,3,6,7}
	InsertSort(array)
	fmt.Println("插入排序：",array)
}
