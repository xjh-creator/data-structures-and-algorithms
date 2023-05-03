package main

import (
	"math"
	"strconv"
)

// PreorderTraversalRecursion 前序遍历,递归法
func PreorderTraversalRecursion(root *Node) []int {
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

// PreorderTraversal 前序遍历,迭代法
func PreorderTraversal(root *Node) []int {
	result := make([]int,0)

	stack := make([]*Node,0)
	stack = append(stack,root)

	for len(stack) > 0{
		e := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		result = append(result,e.value)

		if e.right != nil{
			stack = append(stack,e.right)
		}
		if e.left != nil{
			stack = append(stack,e.left)
		}
	}

	return result
}

// PostorderTraversalRecursion 后序遍历
func PostorderTraversalRecursion(root *Node) []int {
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

// PostorderTraversal 后序遍历
func PostorderTraversal(root *Node) []int {
	result := make([]int,0)

	stack := make([]*Node,0)
	stack = append(stack,root)

	for len(stack) > 0{
		e := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		result = append(result,e.value)

		if e.left != nil{
			stack = append(stack,e.left)
		}
		if e.right != nil{
			stack = append(stack,e.right)
		}
	}

	for left,right:=0,len(result) - 1;right > left;{
		result[left],result[right] = result[right],result[left]
		left++
		right--
	}

	return result
}

// InorderTraversalRecursion 中序遍历
func InorderTraversalRecursion(root *Node) []int {
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

// InorderTraversal 中序遍历
func InorderTraversal(root *Node) []int {
	result := make([]int,0)

	stack := make([]*Node,0)
	cur := root

	for cur != nil || len(stack) > 0{
		if cur != nil{
			stack = append(stack,cur)
			cur = cur.left
		}else{
			cur = stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			result = append(result,cur.value)
			cur = cur.right
		}
	}

	return result
}

// PreorderStandTraversal 前序遍历,统一迭代法
func PreorderStandTraversal(root *Node) []int {
	result := make([]int,0)

	stack := make([]*Node,0)
	stack = append(stack,root)

	for len(stack) > 0{
		e := stack[len(stack) - 1]

		if e != nil{
			stack = stack[:len(stack) - 1]
			if e.right != nil{
				stack = append(stack,e.right)
			}
			if e.left != nil{
				stack = append(stack,e.left)
			}
			stack = append(stack,e)
			stack = append(stack,nil)
		}else{
			stack = stack[:len(stack) - 1]

			node := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			result = append(result,node.value)
		}
	}

	return result
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
// levelOrder 102.二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	ret := make([][]int,0)

	if root == nil{
		return ret
	}

	queue := make([]*TreeNode,0)
	queue = append(queue,root)

	for len(queue) != 0{
		temp := []int{}

		queueLen := len(queue)
		for i:=0;i<queueLen;i++{
			e := queue[0]
			queue = queue[1:]

			temp = append(temp,e.Val)

			if e.Left != nil{
				queue = append(queue,e.Left)
			}
			if e.Right != nil{
				queue = append(queue,e.Right)
			}
		}

		ret = append(ret,temp)
	}

	return ret
}

// invertTree 226.翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	var preOrder func(root *TreeNode)

	preOrder = func(root *TreeNode){
		if root == nil{
			return
		}
		root.Left,root.Right = root.Right,root.Left
		preOrder(root.Left)
		preOrder(root.Right)
	}

	preOrder(root)
	return root
}

// isSymmetric 101. 对称二叉树
func isSymmetric(root *TreeNode) bool {
	var def func(left,right *TreeNode) bool

	def = func(left,right *TreeNode) bool{
		if left == nil && right == nil{
			return true
		}
		if left == nil || right == nil{
			return false
		}
		if left.Val != right.Val{
			return false
		}
		return def(left.Left,right.Right) && def(right.Left,left.Right)
	}

	return def(root.Left,root.Right)
}

// maxDepth 104.二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}

	return max(maxDepth(root.Left),maxDepth(root.Right)) + 1
}

func max(a,b int)int{
	if a > b{
		return a
	}
	return b
}

// minDepth 111.二叉树的最小深度
func minDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}

	if root.Left == nil && root.Right != nil{
		return minDepth(root.Right) + 1
	}

	if root.Left != nil && root.Right == nil{
		return minDepth(root.Left) + 1
	}

	return min(minDepth(root.Left),minDepth(root.Right)) + 1
}

func min(a,b int)int{
	if a < b{
		return a
	}
	return b
}

// countNodes 222.完全二叉树的节点个数
func countNodes(root *TreeNode) int {
	if root == nil{
		return 0
	}

	res := 1
	if root.Left != nil{
		res += countNodes(root.Left)
	}
	if root.Right != nil{
		res += countNodes(root.Right)
	}

	return res
}

// isBalanced 110.平衡二叉树
func isBalanced(root *TreeNode) bool {
	var getDepth func(root *TreeNode) int

	getDepth = func(root *TreeNode)int{
		if root == nil{
			return 0
		}

		lLen,rLen := getDepth(root.Left),getDepth(root.Right)
		if lLen == -1 || rLen == -1{
			return -1
		}
		if lLen - rLen > 1 || rLen - lLen > 1{
			return -1
		}

		if lLen > rLen{
			return lLen + 1
		}

		return rLen + 1
	}

	ret := getDepth(root)
	if ret == -1{
		return false
	}

	return true
}

// binaryTreePaths 257. 二叉树的所有路径
func binaryTreePaths(root *TreeNode) []string {
	ret := []string{}

	var rescur func(root *TreeNode,path string)
	rescur = func(root *TreeNode,path string){
		if root.Left == nil && root.Right == nil{
			path := path + strconv.Itoa(root.Val)
			ret = append(ret,path)
			return
		}

		s := strconv.Itoa(root.Val)
		path = path + s  + "->"
		if root.Left != nil{
			rescur(root.Left,path)
		}
		if root.Right != nil{
			rescur(root.Right,path)
		}
	}

	rescur(root,"")

	return ret
}

// sumOfLeftLeaves 404.左叶子之和
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftValue := sumOfLeftLeaves(root.Left)   // 左

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		leftValue = root.Left.Val             // 中
	}

	rightValue := sumOfLeftLeaves(root.Right) // 右

	return leftValue + rightValue
}

// findBottomLeftValue 513.找树左下角的值
func findBottomLeftValue(root *TreeNode) int {
	ret,depth := 0,0

	var dfs func(root *TreeNode,d int)

	dfs = func(root *TreeNode,d int){
		if root == nil{
			return
		}

		if root.Left == nil && root.Right == nil && d > depth{
			depth = d
			ret = root.Val
		}

		dfs(root.Left,d + 1)
		dfs(root.Right,d + 1)
	}

	dfs(root,1)

	return ret
}

// hasPathSum 112. 路径总和
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	targetSum -= root.Val // 将targetSum在遍历每层的时候都减去本层节点的值
	if root.Left == nil && root.Right == nil && targetSum == 0 { // 如果剩余的targetSum为0, 则正好就是符合的结果
		return true
	}
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum) // 否则递归找
}

// constructMaximumBinaryTree 654.最大二叉树
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	// 找到最大值
	index := findMax(nums)
	// 构造二叉树
	root := &TreeNode {
		Val: nums[index],
		Left: constructMaximumBinaryTree(nums[:index]),   //左半边
		Right: constructMaximumBinaryTree(nums[index+1:]),//右半边
	}
	return root
}
func findMax(nums []int) (index int) {
	for i, v := range nums {
		if nums[index] < v {
			index = i
		}
	}
	return
}

// mergeTrees 617.合并二叉树
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil{
		return root2
	}
	if root2 == nil{
		return root1
	}

	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left,root2.Left)
	root1.Right = mergeTrees(root1.Right,root2.Right)

	return root1
}

// searchBST 700.二叉搜索树中的搜索
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val{
		return root
	}

	if root.Val > val{
		return searchBST(root.Left,val)
	}

	return searchBST(root.Right,val)
}

// isValidBST 98.验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	if root == nil{
		return true
	}

	var check func(root *TreeNode,min,max int) bool
	check = func(root *TreeNode,min,max int) bool{
		if root == nil{
			return true
		}

		if root.Val >= max || root.Val <= min{
			return false
		}

		return check(root.Left,min,root.Val) && check(root.Right,root.Val,max)
	}

	return check(root,int(math.MinInt64),int(math.MaxInt64))
}

// getMinimumDifference 530.二叉搜索树的最小绝对差
func getMinimumDifference(root *TreeNode) int {
	var pre *TreeNode

	min := int(math.MaxInt64)

	var travel func(root *TreeNode)
	travel = func(root *TreeNode){
		if root == nil{
			return
		}

		travel(root.Left)

		if pre != nil && root.Val - pre.Val < min{
			min = root.Val - pre.Val
		}
		pre = root

		travel(root.Right)
	}

	travel(root)

	return min
}

// findMode 501.二叉搜索树中的众数
func findMode(root *TreeNode) []int {
	ret := make([]int,0)

	count := 1
	max := 1
	var pre *TreeNode

	var travel func(root *TreeNode)
	travel = func(root *TreeNode){
		if root == nil{
			return
		}

		travel(root.Left)
		if pre != nil && pre.Val == root.Val {
			count++
		} else {
			count = 1
		}
		if count >= max{
			if count > max && len(ret) > 0{
				ret = []int{root.Val}
			}else{
				ret = append(ret,root.Val)
			}
			max = count
		}

		pre = root
		travel(root.Right)
	}

	travel(root)

	return ret
}