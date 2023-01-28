package linked_list

import "fmt"

func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.next
		cur.next = pre
		pre = cur
		cur = next
	}
	return pre
}

func RcReverseList(head *ListNode) *ListNode {
	return solve(nil,head)
}

func solve(pre,cur *ListNode) *ListNode {
	if cur == nil{
		return pre
	}
	next := cur.next
	cur.next = pre
	pre = cur
	cur = next
	return solve(pre,cur)
}

func PrintListNode(head *ListNode)  {
	for head != nil{
		fmt.Printf("%d-->",head.value.(int))
		head = head.next
	}
}
