package single

import "fmt"

/*
单链表基本操作
*/

// ListNode 链表的节点
type ListNode struct {
	next  *ListNode
	value interface{}
}

// LinkedList 链表
type LinkedList struct {
	head   *ListNode
	length uint
}

// NewListNode 新的链表节点
func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

// GetNext 获取当前节点的下一个节点
func (this *ListNode) GetNext() *ListNode {
	return this.next
}

// GetValue 获取当前节点的值
func (this *ListNode) GetValue() interface{} {
	return this.value
}

// NewLinkedList 新的链表
func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

// InsertAfter 在链表某个节点后面插入节点
func (this *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if nil == p {
		return false
	}
	newNode := NewListNode(v)
	oldNext := p.next
	p.next = newNode
	newNode.next = oldNext
	this.length++
	return true
}

// InsertBefore 在某个节点前面插入节点
func (this *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if nil == p || p == this.head {
		return false
	}
	cur := this.head.next
	pre := this.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}
	newNode := NewListNode(v)
	pre.next = newNode
	newNode.next = cur
	this.length++
	return true
}

// InsertToHead 在链表头部插入节点
func (this *LinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

// InsertToTail 在链表尾部插入节点
func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	for nil != cur.next {
		cur = cur.next
	}
	return this.InsertAfter(cur, v)
}

// FindByIndex 通过索引查找节点
func (this *LinkedList) FindByIndex(index uint) *ListNode {
	if index >= this.length {
		return nil
	}
	cur := this.head.next
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

// DeleteNode 删除传入的节点
func (this *LinkedList) DeleteNode(p *ListNode) bool {
	if nil == p {
		return false
	}
	cur := this.head.next
	pre := this.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}
	pre.next = p.next
	p = nil
	this.length--
	return true
}

// Print 打印链表
func (this *LinkedList) Print() {
	cur := this.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}

// revListed 反转链表
func (this *LinkedList)revListed() *ListNode {
	var pre *ListNode
	cur := this.head
	for cur != nil{
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}
	return cur
}

// checkCircle 链表中环的检测，快慢指针法，存在环总会相遇
func (this *LinkedList)checkCircle() bool {
	if this.head == nil || this.head.next == nil{
		return false
	}
	first := this.head
	second := this.head.next
	for first != nil && second.next != nil{ // 如果是环，first 和 second 都不会为 nil
		first = first.next
		second = second.next.next
		if second == nil{
			return false
		}
		if first == second{
			return true
		}
	}
	return false
}

// mergeSortedLinkedList 合并两个有序链表
func mergeSortedLinkedList(a,b *LinkedList) *LinkedList {
	if a.head == nil{
		return b
	}
	if b.head == nil{
		return a
	}

	tempNode := NewListNode(-1)
	temp := tempNode
	num := 0
	for a.head != nil && b.head != nil{
		if a.head.value.(int) < b.head.value.(int){
			temp.next = a.head
			temp = temp.next
			a.head = a.head.next
		}else{
			temp.next = b.head
			temp = temp.next
			b.head = b.head.next
		}
		num++
	}
	if a.head == nil{
		temp.next = b.head
	}
	if b.head == nil{
		temp.next = a.head
	}
	newList := NewLinkedList()
	newList.head = tempNode.next
	return newList
}

// deleteLastNode 删除倒数第 n 个节点
func (this *LinkedList)deleteLastNode(n int) bool{
	if this.head == nil{
		return false
	}
	low,fast := this.head,this.head
	for i:=0;i < n;i++{ // 快慢指针，中间相距 n 个节点，快指针到链表尾部，慢指针下一个位置就是要删除的节点位置
		if fast.next == nil{
			return false
		}
		fast = fast.next
	}

	for fast.next != nil{
		low = low.next
		fast = fast.next
	}
	low.next = low.next.next
	return true
}

// getMiddleNode 获取链表的中间节点
func (this *LinkedList)getMiddleNode() *ListNode{
	if this.head == nil{
		return nil
	}

	cur,temp := this.head,this.head
	count := 0

	for cur != nil{
		count++
		cur = cur.next
	}
	for i := 0;i < count / 2;i++{
		temp = temp.next
	}
	return temp
}


