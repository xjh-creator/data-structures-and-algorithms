package bfs_dfs

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

/*
	给定一个二叉树，找出其最小深度。最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
*/

// minDepth 给定一个二叉树，找出其最小深度。
func minDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}

	queue := make([]*TreeNode,0)
	queue = append(queue,root)

	depth := 1
	queueLen := len(queue)
	for queueLen != 0{
		sz := queueLen
		for i:=0;i<sz;i++{
			cur := queue[0]
			queue = queue[1:]
			if cur.Left == nil && cur.Right == nil{
				return depth
			}
			if cur.Left != nil{
				queue = append(queue,cur.Left)
			}
			if cur.Right != nil{
				queue = append(queue,cur.Right)
			}
		}
		depth++
		queueLen = len(queue)
	}
	return depth
}

/*
	你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。每个拨轮可以自由旋转：例如把 '9' 变为 '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。

	锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。

	列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。

	字符串 target 代表可以解锁的数字，你需要给出解锁需要的最小旋转次数，如果无论如何不能解锁，返回 -1 。
*/
// OpenLock 打开转盘锁
func OpenLock(deadends []string, target string) int {
	queue := make([]string,0)
	queue = append(queue,"0000")

	dead := make(map[string]bool)
	for _,v := range deadends{
		dead[v] = true
	}
	visited := make(map[string]bool)
	minStep := 0
	queueLen := len(queue)
	for queueLen != 0{
		sz := queueLen
		for i:=0;i<sz;i++{
			top := queue[0]
			queue = queue[1:]
			if dead[top]{
				continue
			}
			if top == target{
				return minStep
			}
			for j:=0;j<4;j++{
				up := plusOne(top,j)
				if !visited[up]{
					queue = append(queue,up)
					visited[up] = true
				}
				down := minusOne(top,j)
				if !visited[down]{
					queue = append(queue,down)
					visited[down] = true
				}
			}
		}
		minStep++
		queueLen = len(queue)
	}
	return -1
}

// 将 s[j] 向上拨动一次
func plusOne(s string,j int) string {
	ch := []byte(s)
    if ch[j] == '9'{
    	ch[j] = '0'
	} else{
		ch[j] += 1
	}
	return string(ch)
}

// 将 s[i] 向下拨动一次
func minusOne(s string,j int) string {
	ch := []byte(s)
	if ch[j] == '0' {
		ch[j] = '9'
	} else {
		ch[j] += 1
	}
	return string(ch)
}