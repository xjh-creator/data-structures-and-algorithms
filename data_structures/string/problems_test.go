package string

import (
	"fmt"
	"testing"
)

func TestReverseStr(t *testing.T)  {
	s := "abcdefg"
	k := 2
	fmt.Println(ReverseStr(s,k))
}

func TestReverseWords(t *testing.T)  {
	s := "a good   example"
	fmt.Println(ReverseWords(s))
}

func TestReverseLeftWords(t *testing.T)  {
	s := "lrloseumgh"
	fmt.Println(ReverseLeftWords(s,6))
}

func TestRepeatedSubstringPattern(t *testing.T)  {
	s := "aba"
	fmt.Println(RepeatedSubstringPattern(s))
}

func TestStrStr(t *testing.T)  {
	haystack := "mississippi"
	needle :=       "issipi"

	fmt.Println(StrStr(haystack,needle))
}


