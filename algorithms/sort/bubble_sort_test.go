package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T)  {
	array := []int{1,4,5,8,9,2,10,3,6,7}
	BubbleSort(array)
	fmt.Println("冒泡排序：",array)
}
