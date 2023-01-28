package main

// PreorderTraversal 前序遍历,递归法
func PreorderTraversal(root *Node) []int {
	result := make([]int,0)
	var travelTree func(node *Node)
	travelTree = func(node *Node){
		if node == nil{
			return
		}
		result = append(result,node.value)
		travelTree(node.left)
		travelTree(node.right)
	}
	travelTree(root)
	return result
}

// PostorderTraversal 后序遍历
func PostorderTraversal(root *Node) []int {
	result := make([]int,0)
	var travelTree func(node *Node)
	travelTree = func(node *Node){
		if node == nil{
			return
		}
		travelTree(node.left)
		travelTree(node.right)
		result = append(result,node.value)
	}
	travelTree(root)
	return result
}

// InorderTraversal 中序遍历
func InorderTraversal(root *Node) []int {
	result := []int{}
	var travelTree func(node *Node)
	travelTree = func(node *Node) {
		if node == nil{
			return
		}
		travelTree(node.left)
		result = append(result,node.value)
		travelTree(node.right)
	}
	travelTree(root)
	return result
}