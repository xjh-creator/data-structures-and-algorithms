package main

import "fmt"

type binarySearchTreeNode struct {
	value int
	left  *binarySearchTreeNode
	right *binarySearchTreeNode
}

// find 二叉搜索树查找
func (root *binarySearchTreeNode) find(data int) *binarySearchTreeNode {
	node := root
	for node != nil {
		if node.value > data {
			node = node.left
		} else if node.value < data {
			node = node.right
		} else {
			fmt.Println("找到了")
			return node
		}
	}
	return nil
}

// insert 二叉搜索树插入
func (root *binarySearchTreeNode) insert(data int) {
	node := root
	if node == nil {
		root = &binarySearchTreeNode{value: data}
	}
	for node != nil {
		if data > node.value {
			if node.right == nil {
				node.right = &binarySearchTreeNode{value: data}
				return
			}
			node = node.right
		} else { //data < node.value
			if node.left == nil {
				node.left = &binarySearchTreeNode{value: data}
				return
			}
			node = node.left
		}
	}
}

// delete 二叉搜索树删除
func (root *binarySearchTreeNode) delete(data int) {
	node := root //node指向要删除的节点
	nodeFather := &binarySearchTreeNode{}
	for node != nil && node.value != data {
		nodeFather = node
		if data > node.value {
			node = node.right
		} else {
			node = node.left
		}
	}
	if node == nil {
		return //没有找到
	}
	//两个子节点都不为空
	if node.left != nil && node.right != nil { //查找右子树中最小的点
		minNode := node.right
		minNodeP := node //右子树中最小的点的父节点
		for minNode.left != nil {
			minNodeP = minNode
			minNode = minNode.left
		}
		node.value = minNode.value //将最小节点的值替换到要删除的节点中
		node = minNode //下面就是删除 minNode
		nodeFather = minNodeP
	}

	// 删除的节点是叶子节点或者仅有一个子节点
	child := &binarySearchTreeNode{} //node的子节点
	if node.left != nil {
		child = node.left
	} else if node.right != nil {
		child = node.right
	} else {
		child = nil
	}
	if nodeFather == nil {  // 删除的是根节点
		root = child
	} else if nodeFather.left == node {
		nodeFather.left = child
	} else {
		nodeFather.right = child
	}

}

func main() {
	root := &binarySearchTreeNode{value: 13}
	node1 := &binarySearchTreeNode{value: 10}
	node2 := &binarySearchTreeNode{value: 16}
	node3 := &binarySearchTreeNode{value: 9}
	node4 := &binarySearchTreeNode{value: 11}
	node5 := &binarySearchTreeNode{value: 14}
	root.left = node1
	root.right = node2
	node1.left = node3
	node1.right = node4
	node2.left = node5
	result := root.find(11)
	fmt.Println(result)
}
