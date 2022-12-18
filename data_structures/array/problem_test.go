package array

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	nums := []int{-1,0,3,5,9,12}
	target := 9
	fmt.Println(Search(nums,target))
}

func TestRemoveElement(t *testing.T) {
	nums1 := []int{0,1,2,2,3,0,4,2}
	nums2 := []int{3,2,2,3}
	result1 := RemoveElement(nums1,2)
	fmt.Println(result1)
	fmt.Sprintf("num1 的count:%d\n",result1)
	result2 := RemoveElement(nums2,2)
	fmt.Println(result2)
	fmt.Sprintf("num2 的count:%d\n",result2)
}

func TestSortedSquares(t *testing.T) {
	resutl1 := SortedSquares([]int{-4,-1,0,3,10})
	fmt.Println(resutl1)
	resutl2 := SortedSquares([]int{-7,-3,2,3,11})
	fmt.Println(resutl2)
}

