package test

import (
	"fmt"
	"testing"
)

func TestCommonChars(t *testing.T)  {
	w := []string{"bella","label","roller"}

	ret := CommonChars(w)
	fmt.Println(ret)
}
