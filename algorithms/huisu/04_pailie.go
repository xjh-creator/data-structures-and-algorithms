package main

import "sort"

/*
	给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

    输入: [1,2,3]
    输出: [ [1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1] ]

*/
// Permute 46.全排列
func Permute(nums []int) [][]int {
	res := make([][]int,0)
	path := make([]int,0)
	numsLen := len(nums)
	state := make([]bool,numsLen)

	var dfs func()
	dfs = func() {
		pathLen := len(path)
		if pathLen == numsLen{
			temp := make([]int,pathLen)
			copy(temp,path)
			res = append(res,temp)
			return
		}

		for i:=0;i<numsLen;i++{
			if !state[i]{
				path = append(path,nums[i])
				state[i] = true
				dfs()
				state[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	dfs()
	return res
}

/*
	给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

	示例 1：

    输入：nums = [1,1,2]
    输出： [[1,1,2], [1,2,1], [2,1,1]]

	示例 2：

    输入：nums = [1,2,3]
    输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

	提示：

    1 <= nums.length <= 8
    -10 <= nums[i] <= 10

*/
// PermuteUnique
func PermuteUnique(nums []int) [][]int {
	res := make([][]int,0)
	path := make([]int,0)
	numsLen := len(nums)
	state := make([]bool,numsLen)
	sort.Ints(nums)

	var dfs func()
	dfs = func() {
		pathLen := len(path)
		if pathLen == numsLen{
			temp := make([]int,pathLen)
			copy(temp,path)
			res = append(res,temp)
			return
		}

		for i:=0;i<numsLen;i++{
			if i != 0 && nums[i] == nums[i-1] && !state[i-1]{
				continue
			}

			if !state[i]{
				path = append(path,nums[i])
				state[i] = true
				dfs()
				state[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	dfs()
	return res
}
