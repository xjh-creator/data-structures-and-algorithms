package linked_list

import "fmt"

func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil{
		temp := cur.next
		cur.next = pre
		pre = cur
		pre = temp
	}
	return cur
}

func PrintListNode(head *ListNode)  {
	for head.next != nil{
		fmt.Printf("%d-->",head.value.(int))
		head = head.next
	}
}
