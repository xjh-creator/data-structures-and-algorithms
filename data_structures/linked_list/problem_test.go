package linked_list

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	head := NewListNode(1)
	oneNode := NewListNode(2)
	twoNode := NewListNode(3)
	threeNode := NewListNode(4)
	head.next = oneNode
	oneNode.next = twoNode
	twoNode.next = threeNode
	PrintListNode(head)
    fmt.Println()
	result := ReverseList(head)
	PrintListNode(result)
}

func TestRcReverseList(t *testing.T) {
	head := NewListNode(1)
	oneNode := NewListNode(2)
	twoNode := NewListNode(3)
	threeNode := NewListNode(4)
	head.next = oneNode
	oneNode.next = twoNode
	twoNode.next = threeNode
	PrintListNode(head)
	fmt.Println()
	result := RcReverseList(head)
	PrintListNode(result)
}