package array

//给定一个n个元素有序的（升序）整型数组 nums 和一个目标值 target  ，
//写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

// Search 二分查找
func Search(nums []int,target int) int {
	length := len(nums)
	left,right := 0,length - 1

	for left < right{
		middle := left + (right - left) / 2
		if nums[middle] == target{
			return middle
		}else if nums[middle] > target{
			right = middle - 1
		}else{
			left = middle + 1
		}
	}
	return -1
}

/*
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于val的元素，并返回移除后数组的新长度。

不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并原地修改输入数组。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

示例1: 给定 nums = [3,2,2,3], val = 3, 函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。 你不需要考虑数组中超出新长度后面的元素。

示例2: 给定 nums = [0,1,2,2,3,0,4,2], val = 2, 函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。
*/

// RemoveElement 移除元素
func RemoveElement(nums []int, val int) int {
	numsLen := len(nums)
	res := 0
	for i:=0;i<numsLen;i++{
		if nums[i] != val{
			nums[res] = nums[i]
			res++
		}
	}
	nums = nums[:res]
	return res
}

/*
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

示例 1： 输入：nums = [-4,-1,0,3,10] 输出：[0,1,9,16,100] 解释：平方后，数组变为 [16,1,0,9,100]，排序后，数组变为 [0,1,9,16,100]

示例 2： 输入：nums = [-7,-3,2,3,11] 输出：[4,9,9,49,121]
*/

// SortedSquares 有序数组的平方
func SortedSquares(nums []int) []int {
	numsLen := len(nums)
	l,r,k := 0,numsLen - 1,numsLen -1
	result := make([]int,numsLen)
	for l <= r{
		lm,rm := nums[l] * nums[l],nums[r] * nums[r]
		if lm <= rm{
			result[k] = rm
			r--
		}else{
			result[k] = lm
			l++
		}
		k--
	}
	return result
}
