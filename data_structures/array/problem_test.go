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

