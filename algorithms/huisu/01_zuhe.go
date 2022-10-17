package main

// 问题：给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
/*
	示例:
	输入: n = 4, k = 2
	输出:
	[
		[2,4],
		[3,4],
		[2,3],
		[1,2],
		[1,3],
		[1,4],
	]
*/

var res [][]int
func combine(n int,k int) [][]int {
	res=[][]int{}
	if n <= 0 || k <= 0 || k > n {
		return res
	}
	backtracking1(n, k, 1, []int{})
	return res
}

// backtracking1
func backtracking1(n,k,start int,track []int){
	if len(track)==k{
		temp:=make([]int,k)
		copy(temp,track)
		res=append(res,temp)
	}
	if len(track)+n-start+1 < k { // ???
		return
	}
	for i:=start;i<=n;i++{
		track=append(track,i) // 处理节点
		backtracking1(n,k,i+1,track) // 递归
		track=track[:len(track)-1] // 回溯，撤销处理的节点
	}
}
