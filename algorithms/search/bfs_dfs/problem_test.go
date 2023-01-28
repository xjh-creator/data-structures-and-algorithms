package bfs_dfs

import (
	"fmt"
	"testing"
)

func TestOpenLock(t *testing.T) {
	deadends := []string{"0201","0101","0102","1212","2002"}
	str := "0202"

	result := OpenLock(deadends,str)
	fmt.Println(result)

}
