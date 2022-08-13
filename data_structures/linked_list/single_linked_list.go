package main

import "fmt"

type listNode struct {
	val int
	next *listNode
}

//insert 在链表尾部插入值为v的新节点
func (n *listNode)insert(v int)  {
	dummyHead := &listNode{}
	dummyHead.next = n
	cur := dummyHead
	for cur.next != nil{
		cur = cur.next
	}
	cur.next = &listNode{val:v}
	fmt.Println("插入了数值",v)
}

//delete 从列表中删除值等于v的节点
func (n *listNode)delete(v int) {
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
}

func (n *listNode)show()  {
	head := n
	for i := 1;head != nil;i++{
		fmt.Printf("第%d个节点的值为：%d \n",i,head.val)
		head = head.next
	}
}

//get 获取链表中第index个节点的值
func (n *listNode)get(index int)  {
	head := n
	for i := 1;head != nil;i++{
		if i == index{
			fmt.Printf("第%d个节点的值为：%d \n",i,head.val)
			return
		}
		head = head.next
	}
}

//addAtHead 在链表的第一个元素之前添加一个值为 val 的节点。
func (n *listNode)addAtHead(val int){
	newNode := &listNode{
		val: val,
		next: n,
	}
	n = newNode
}


func main()  {
	head := &listNode{val: 1}
	node1 := &listNode{val: 2}
	node2 := &listNode{val: 3}
	head.next = node1
	node1.next = node2
	//head.show()
    //head.insert(4)
	//head.show()
	//head.delete(2)
	//head.get(2)
	head.show()
	head.addAtHead(0)
	head.show()
}