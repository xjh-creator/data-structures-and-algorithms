package stack

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	s := NewArrayStack()
	str := "(({{}}))[]"
	fmt.Println("测试结果:",s.IsValid(str))
	s.Print()
}

func TestRemoveDuplicates(t *testing.T) {
	s := NewArrayStack()
	result := s.RemoveDuplicates("abbaca")
	fmt.Println(result)
}