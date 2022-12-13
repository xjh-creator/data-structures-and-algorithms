package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

// preOrder 前序遍历
func preOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.value)
	preOrder(root.left)
	preOrder(root.right)
}

// inOrder 中序遍历，
func inOrder(root *Node) {
	if root == nil {
		return
	}
	inOrder(root.left)
	fmt.Println(root.value)
	inOrder(root.right)
}

// postOrder 后序遍历
func postOrder(root *Node) {
	if root == nil {
		return
	}
	postOrder(root.left)
	postOrder(root.right)
	fmt.Println(root.value)
}

// levelOrder 层序遍历
//用一个队列来存储
//思路：每次队列存储二叉树的一层，然后判断存储的节点有没有子节点，把上一层的节点一边弹出
//队列一边把下一层的节点加入队列，弹出的节点先存入一维数组，再存入二维数组
func levelOrder(root *Node) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	//初始化队列
	queue := NewQueue()
	//先把跟节点存入队列
	queue.Push(root)

	//temp存放每一层的节点的值
	temp := make([]int, 0)
	for queue.Len() > 0 {
		//length为当前队列所存放那一层的节点数
		length := queue.Len()
		//遍历每一层的节点
		for i := 0; i < length; i++ {
			//node弹出的节点
			node := queue.Pop()
			//判断弹出的节点有没有子节点，有的话加入
			if node.left != nil {
				queue.Push(node.left)
			}
			if node.right != nil {
				queue.Push(node.right)
			}
			temp = append(temp, node.value)
		}
		result = append(result, temp)
		temp = []int{}
	}
	return result
}

func main() {
	root := &Node{value: 10}
	node1 := &Node{value: 32}
	node2 := &Node{value: 23}
	node3 := &Node{value: 21}
	node4 := &Node{value: 56}
	node5 := &Node{value: 34}
	root.left = node1
	root.right = node2
	node1.left = node3
	node1.right = node4
	node2.left = node5
	postOrder(root)
}

type Queue struct {
	queue []*Node
}

func NewQueue() *Queue {
	return &Queue{queue: make([]*Node, 0)}
}

func (q *Queue) Push(val *Node) {
	q.queue = append(q.queue, val)
}

func (q *Queue) Pop() *Node {
	if len(q.queue) == 0 {
		return nil
	}
	result := q.queue[0]
	q.queue = q.queue[1:]
	return result
}

func (q *Queue) Len() int {
	return len(q.queue)
}
