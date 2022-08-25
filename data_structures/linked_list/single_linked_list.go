package main

import "fmt"

type listNode struct {
	val int
	next *listNode
}

//insert 在链表尾部插入值为v的新节点
func (n *listNode)insert(v int) *listNode {
	dummyHead := &listNode{}
	dummyHead.next = n
	cur := dummyHead
	for cur.next != nil{
		cur = cur.next
	}
	cur.next = &listNode{val:v}
	fmt.Println("插入了数值",v)
	return dummyHead.next
}

//delete 从列表中删除值等于v的节点
func (n *listNode)delete(v int) *listNode{
	dummyHead := &listNode{}
	dummyHead.next = n
	cur := dummyHead
	for cur.next.val != v{
		cur = cur.next
	}
	if cur.next.val == v{
		cur.next = cur.next.next
	}
	fmt.Println("删除了数值",v)
	return dummyHead.next
}

func (n *listNode)show()  {
	head := n
	for i := 1;head != nil;i++{
		fmt.Printf("第%d个节点的值为：%d \n",i,head.val)
		head = head.next
	}
}

//get 获取链表中第index个节点的值
func (n *listNode)get(index int) int {
	head := n
	for i := 1;head != nil;i++{
		if i == index{
			fmt.Printf("第%d个节点的值为：%d \n",i,head.val)
			return head.val
		}
		head = head.next
	}
	return 0
}

//addAtHead 在链表的第一个元素之前添加一个值为 val 的节点。
func (n *listNode)addAtHead(val int) *listNode{
	return &listNode{
		val: val,
		next: n,
	}
}

// revListed 反转链表
func revListed(head *listNode) *listNode {
	var pre *listNode
	cur := head
	for cur != nil{
		next := cur.next
		cur.next = pre
		pre = cur
		cur = next
	}
	return pre
}

// changeNodePerTwo 两两交换链表中相邻的结点
func changeNodePerTwo(head *listNode) *listNode {
	dummyHead := &listNode{}
	dummyHead.next = head
	cur := dummyHead

	for head != nil && head.next != nil{
		cur.next = head.next
		next := head.next.next
		head.next.next = head
		head.next = next

		cur = head
		head = next
	}
	return dummyHead.next
}

//delLaterNode 删除最后第n个节点
func delLaterNode(head *listNode,n int)*listNode{
	dummyHead := &listNode{next: head}
	pre := dummyHead
	cur := head

	i := 1
	if cur != nil{
		cur = cur.next
		if i > n{
			pre = pre.next
		}
		i++
	}
	pre.next = pre.next.next
	return dummyHead.next
}

func getIntersectionNode(headA,headB *listNode) *listNode {
	curA := headA
	curB := headB
	lenA, lenB := 0, 0
	// 求A，B的长度
	for curA != nil {
		curA = curA.next
		lenA++
	}
	for curB != nil {
		curB = curB.next
		lenB++
	}
	var step int
	var fast, slow *listNode
	// 请求长度差，并且让更长的链表先走相差的长度
	if lenA > lenB {
		step = lenA - lenB
		fast, slow = headA, headB
	} else {
		step = lenB - lenA
		fast, slow = headB, headA
	}
	for i:=0; i < step; i++ {
		fast = fast.next
	}
	// 遍历两个链表遇到相同则跳出遍历
	for fast != slow {
		fast = fast.next
		slow = slow.next
	}
	return fast
}

func main()  {
	head1 := &listNode{val: 1}
	node1 := &listNode{val: 2}
	node2 := &listNode{val: 3}
	node3 := &listNode{val: 4}
	node4 := &listNode{val: 3}
	node5 := &listNode{val: 4}
	node6 := &listNode{val: 10}
	head1.next = node1
	node1.next = node2
	node2.next = node3
	//changeNodePerTwo(head).show()
	head2 := &listNode{val: 1}
	head2.next = node6
	node6.next = node4
	node4.next = node5

	getIntersectionNode(head1,head2).show()

	//delLaterNode(head,1).show()
}

//func main()  {
//	head := &listNode{val: 1}
//	node1 := &listNode{val: 2}
//	node2 := &listNode{val: 3}
//	head.next = node1
//	node1.next = node2
//	//head.show()
//    //head.insert(4)
//	//head.show()
//	//head.delete(2)
//	//head.get(2)
//	head.show()
//	head.addAtHead(0)
//	head.show()
//}