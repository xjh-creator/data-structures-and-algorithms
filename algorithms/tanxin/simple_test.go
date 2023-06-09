package tanxin

import (
	"fmt"
	"testing"
)

func TestFindContentChildren(t *testing.T){
	g := []int{1,2,3}
	s := []int{3}
	fmt.Println(FindContentChildren(g,s))
}

func TestJump(t *testing.T){
	ret := largestSumAfterKNegations([]int{4,2,3},1)

	t.Log(ret)
}
