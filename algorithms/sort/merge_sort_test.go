package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T)  {
	array := []int{1,4,5,8,9,2,10,3,6,7}
	MergeSort(array)
	fmt.Println("---",array)
	//for _,v := range array{
	//	fmt.Println(v)
	//}
}
