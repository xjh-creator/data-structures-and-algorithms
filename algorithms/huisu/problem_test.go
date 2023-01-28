package main

/*
	给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
*/
// Permute 全排列
func Permute(nums []int) [][]int {
	var (
		res [][]int
		path  []int
		st    []bool   // state的缩写
	)

	res, path = make([][]int, 0), make([]int, 0, len(nums))
	st = make([]bool, len(nums))

	var permute func(nums []int, cur int)
	permute = func(nums []int, cur int) {
		if cur == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		for i := 0; i < len(nums); i++ {
			if !st[i] {
				path = append(path, nums[i])
				st[i] = true
				permute(nums, cur + 1)
				st[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	permute(nums,0)
	return res
}

