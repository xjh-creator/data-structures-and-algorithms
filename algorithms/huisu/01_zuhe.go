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

func combine(n int,k int) [][]int {
	res := [][]int{}
	if n <= 0 || k <= 0 || k > n {
		return res
	}
	var backtracking1 func(n,k,start int,track []int)
	backtracking1 = func (n,k,start int,track []int){
		if len(track)==k{
			temp:=make([]int,k)
			copy(temp,track)
			res=append(res,temp)
		}
		if len(track)+n-start+1 < k { // 优化点：走过的 + 未走的 < 限定的 需要跳过
			return
		}
		for i:=start;i<=n;i++{
			track=append(track,i) // 处理节点
			backtracking1(n,k,i+1,track) // 递归
			track=track[:len(track)-1] // 回溯，撤销处理的节点
		}
	}
	backtracking1(n, k, 1, []int{})
	return res
}

/*
	找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。
说明：

    所有数字都是正整数。
    解集不能包含重复的组合。

    示例 1: 输入: k = 3, n = 7 输出: [[1,2,4]]

    示例 2: 输入: k = 3, n = 9 输出: [[1,2,6], [1,3,5], [2,3,4]]
*/
// combinationSum3 216.组合总和III
func combinationSum3(k int, n int) [][]int {
	var (
		res [][]int
		path  []int
	)

	var dfs func(k, n int, start int, sum int)
	dfs = func(k, n int, start int, sum int){
		if len(path) == k {
			if sum == n {
				tmp := make([]int, k)
				copy(tmp, path)
				res = append(res, tmp)
			}
			return
		}
		for i := start; i <= 9; i++ {
			if sum + i > n || 9-i+1 < k-len(path) {
				break
			}
			path = append(path, i)
			dfs(k, n, i+1, sum+i)
			path = path[:len(path)-1]
		}
	}

	res, path = make([][]int, 0), make([]int, 0, k)
	dfs(k, n, 1, 0)
	return res
}

/*
	给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

	给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

	示例: 输入："23" 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].

	说明：尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
*/
// LetterCombinations 电话号码的字母组合
func LetterCombinations(digits string) []string {
	if digits == ""{
		return nil
	}
	path1,res := make([]rune,0),make([]string,0)

	digitsMap := map[byte]string{
		'2':"abc",
		'3':"def",
		'4':"ghi",
		'5':"jkl",
		'6':"mno",
		'7':"pqrs",
		'8':"tuv",
		'9':"wxyz",
	}

	var dfs func(digits string,start int,path []rune)
	dfs = func(digits string,start int,path []rune) {
		if len(path) == len(digits){
			temp := string(path)
			res = append(res,temp)
			return
		}

		for _,v := range digitsMap[digits[start]]{
			path = append(path,v)
			dfs(digits,start+1,path)
			path = path[:len(path)-1]
		}
	}
	dfs(digits,0,path1)
	return res
}

