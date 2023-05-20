package main

import "sort"

/*
	给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

	说明：解集不能包含重复的子集。

	示例: 输入: nums = [1,2,3] 输出: [ [3],   [1],   [2],   [1,2,3],   [1,3],   [2,3],   [1,2],   [] ]
*/
// subsets 78.子集
func subsets(nums []int) [][]int {
	res := make([][]int,0)
	path := make([]int,0,len(nums))

	var dfs func(start int)
	dfs = func(start int){
		temp := make([]int,len(path))
		copy(temp,path)
		res = append(res,temp)

		for i:=start;i<len(nums);i++{
			path = append(path,nums[i])
			dfs(i+1)
			path = path[:len(path) - 1]
		}
	}

	dfs(0)

	return res
}

/*
	给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

	说明：解集不能包含重复的子集。

	示例:

    输入: [1,2,2]
    输出: [ [2], [1], [1,2,2], [2,2], [1,2], [] ]

*/
// subsetsWithDup 90.子集II
func subsetsWithDup(nums []int) [][]int {
	res := make([][]int,0)
	length := len(nums)
	path := make([]int,0,length)

	sort.Ints(nums)
	var dfs func(start int)
	dfs = func(start int) {
		temp := make([]int,len(path))
		copy(temp,path)
		res = append(res,temp)

		for i:=start;i<length;i++{
			if i != start && nums[i] == nums[i-1]{
				continue
			}
			path = append(path,nums[i])
			dfs(i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}

/*
	给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是2。

示例:

    输入: [4, 6, 7, 7]
    输出: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7], [4,7,7]]

	说明:

    给定数组的长度不会超过15。
    数组中的整数范围是 [-100,100]。
    给定数组中可能包含重复数字，相等的数字应该被视为递增的一种情况。

*/
// FindSubsequences 491.递增子序列
func FindSubsequences(nums []int) [][]int {
	res := make([][]int,0)
	path := make([]int,0)
	numsLen := len(nums)

	var dfs func(start int)
	dfs = func(start int) {
		pathLen := len(path)
		if pathLen >= 2{
			temp := make([]int,len(path))
			copy(temp,path)
			res = append(res,temp)
		}

		tempMap := map[int]bool{}
		for i:=start;i<numsLen;i++{
			if tempMap[nums[i]]{
				continue
			}
			if pathLen == 0 || nums[i] >= path[pathLen - 1]{
				tempMap[nums[i]] = true
				path = append(path,nums[i])
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return res
}