package main

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T){
	fmt.Println(Permute([]int{1,2,3}))
}

func TestPermuteUnique(t *testing.T){
	fmt.Println(PermuteUnique([]int{1,1,2}))
}


