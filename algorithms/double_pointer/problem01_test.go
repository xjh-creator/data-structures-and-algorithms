package double_pointer

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T)  {
	temp := []int{2,3,2,3}
	res := removeElement(temp,3)
    fmt.Println(res,"--",temp)
}
