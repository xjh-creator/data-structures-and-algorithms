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
