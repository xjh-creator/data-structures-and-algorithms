package main

import (
	"fmt"
	"testing"
)

func TestPartition(t *testing.T)  {
	result := Partition("aab")
	fmt.Println(result)
}

func TestRestoreIpAddresses(t *testing.T)  {
	result := RestoreIpAddresses("25525511135")
	fmt.Println(result)
}
