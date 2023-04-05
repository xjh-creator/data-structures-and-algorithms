package linked_list

import "fmt"

/*
	给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点
*/

// removeElements 203 移除列表的元素
func removeElements(head *ListNode, val int) *ListNode {
	temp := &ListNode{next:head}
	for cur:=temp;cur.next != nil;{
		if cur.next.value == val{
			cur.next = cur.next.next
		}else{
			cur = cur.next
		}
	}
	return temp.next
}

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

/**
	实现 MyLinkedList 类：

    MyLinkedList() 初始化 MyLinkedList 对象。
    int get(int index) 获取链表中下标为 index 的节点的值。如果下标无效，则返回 -1 。
    void addAtHead(int val) 将一个值为 val 的节点插入到链表中第一个元素之前。在插入完成后，新节点会成为链表的第一个节点。
    void addAtTail(int val) 将一个值为 val 的节点追加到链表中作为链表的最后一个元素。
    void addAtIndex(int index, int val) 将一个值为 val 的节点插入到链表中下标为 index 的节点之前。如果 index 等于链表的长度，那么该节点会被追加到链表的末尾。如果 index 比长度更大，该节点将 不会插入 到链表中。
    void deleteAtIndex(int index) 如果下标有效，则删除链表中下标为 index 的节点。

**/

// MyLinkedList 707.设计链表
type MyLinkedList struct {
	head *Node
	size int
}

type Node struct{
	val int
	next *Node
}


func Constructor() MyLinkedList {
	return MyLinkedList{
		head:&Node{},
		size:0,
	}
}


func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	cur := this.head
	for i := 0; i <= index; i++ {
		cur = cur.next
	}
	return cur.val
}


func (this *MyLinkedList) AddAtHead(val int)  {
	this.AddAtIndex(0, val)
}


func (this *MyLinkedList) AddAtTail(val int)  {
	this.AddAtIndex(this.size, val)
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index > this.size {
		return
	}

	this.size++
	pred := this.head
	for i := 0; i < index; i++ {
		pred = pred.next
	}
	toAdd := &Node{val, pred.next}
	pred.next = toAdd
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index < 0 || index >= this.size {
		return
	}
	this.size--
	pred := this.head
	for i := 0; i < index; i++ {
		pred = pred.next
	}
	pred.next = pred.next.next
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

/*
	题意：反转一个单链表。

	示例: 输入: 1->2->3->4->5->NULL 输出: 5->4->3->2->1->NULL
*/
// reverseList 206.反转链表
func reverseList(head *ListNode) *ListNode {
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
	return solve(cur,next)
}

func PrintListNode(head *ListNode)  {
	for head != nil{
		fmt.Printf("%d-->",head.value.(int))
		head = head.next
	}
}

/*
	给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。
	你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
*/

// swapPairs 24. 两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {
	temp := &ListNode{
		next:head,
	}

	pre := temp
	for head != nil && head.next != nil{
		next := head.next.next
		pre.next = head.next
		head.next.next = head
		head.next = next

		pre = head
		head = head.next
	}

	return temp.next
}

/*
	给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
*/

// removeNthFromEnd 19. 删除链表的倒数第 N 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre := &ListNode{
		next:head,
	}

	cur := head
	for i:=0;i<n;i++{
		cur = cur.next
	}

	// 删除链表只有一个值且删除倒数第一个节点的值
	if cur == nil && n == 1{
		return nil
	}

	step := 0 // 计算 pre 移动的次数
	for cur != nil{
		cur = cur.next
		pre = pre.next
		step++
	}

	pre.next = pre.next.next

	if step == 0{ // pre 没有移动，则返回其下一个节点，即新头节点
		return pre.next
	}

	return head
}
