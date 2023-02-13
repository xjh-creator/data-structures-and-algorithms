package main

import (
	"fmt"
	"testing"
)

func TestFindSubsequences(t *testing.T)  {
	result := FindSubsequences([]int{4,6,7,7})
	fmt.Println(result)
}
