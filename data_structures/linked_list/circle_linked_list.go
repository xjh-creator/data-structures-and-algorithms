package linked_list

import "fmt"

type circleNode struct {
	ID   int
	Next *circleNode
}

//Delete 删除节点 , 分非队尾与队尾两种情况去删除
func (head *circleNode) Delete(deleteNode *circleNode) *circleNode {
	//1.找到头结点
	temp := head
	helper := head //指向环形节点的最后一个
	//空链表
	if temp.Next == nil {
		return head
	}
	for {
		if helper.Next == head {
			break
		}
		helper = helper.Next
	}
	flag := true
	for {
		if temp.Next == head {
			break
		}
		if temp.ID == deleteNode.ID {
			if temp == head {
				head = head.Next
			}
			helper.Next = temp.Next
			flag = false
			break
		}
		temp = temp.Next
		helper = helper.Next //移动，一旦找到就要删除的节点
	}
	if flag {
		if temp.ID == deleteNode.ID {
			helper.Next = temp.Next
			fmt.Printf("[%d]\n", deleteNode.ID)
		} else {
			fmt.Println("没有找到")
		}
	}
	return head
}

//Insert 添加新节点
func (head *circleNode) Insert(newNode *circleNode){
	//判断是否插入第一个节点
	if head.Next == nil {
		head.ID = newNode.ID
		return
	}
	//1.找到头结点
	temp := head
	//2.让指针移动到尾部
	for {
		if temp.Next == head {
			break
		}
		temp = temp.Next
	}
	//3.把新节点添加上去
	temp.Next = newNode
	newNode.Next = head
}

func (head *circleNode) Show() {
	temp := head
	if temp.Next == nil { //循环队列，节点的下一个节点都不为空
		return
	}
	for {
		fmt.Printf("[%d]==>", temp.ID)
		if temp.Next == head {
			break
		}
		temp = temp.Next
	}
}