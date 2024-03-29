package hash

import (
	"fmt"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	str1 := "anagram"
	str2 := "nagaram"
	fmt.Printf("%s和%s是否是异同字符串: ",str1,str2)
	fmt.Println(IsAnagram(str1,str2))

	str3 := "rat"
	str4 := "car"
	fmt.Printf("%s和%s是否是异同字符串: ",str3,str4)
	fmt.Println(IsAnagram(str3,str4))
}

func TestFourSum(t *testing.T) {
	nums := []int{1,0,-1,0,-2,2}

	ret := fourSum(nums,0)
	fmt.Println(ret)

}
