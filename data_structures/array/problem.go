package array

import "sort"

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

/*
给定一个含有n个正整数的数组和一个正整数s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。

示例：

输入：s = 7, nums = [2,3,1,2,4,3] 输出：2 解释：子数组[4,3]是该条件下的长度最小的子数组。
*/

// MinSubArrayLen 长度最小子数组
func MinSubArrayLen(target int, nums []int) int {
	i := 0
	l := len(nums)  // 数组长度
	sum := 0        // 子数组之和
	result := l + 1 // 初始化返回长度为l+1，目的是为了判断“不存在符合条件的子数组，返回0”的情况
	for j := 0; j < l; j++ {
		sum += nums[j]
		for sum >= target {
			subLength := j - i + 1
			if subLength < result {
				result = subLength
			}
			sum -= nums[i]
			i++
		}
	}
	if result == l+1 {
		return 0
	} else {
		return result
	}
}

/*
给定一个正整数n，生成一个包含 1 到n^2所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例:

输入: 3 输出: [ [ 1, 2, 3 ], [ 8, 9, 4 ], [ 7, 6, 5 ] ]
*/
func funcs()  {

}

/*
	给定一个非空的整数数组，返回其中出现频率前 k 高的元素。

示例 1:

    输入: nums = [1,1,1,2,2,3], k = 2
    输出: [1,2]

示例 2:

    输入: nums = [1], k = 1
    输出: [1]

提示：

    你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
    你的算法的时间复杂度必须优于 O(nlogn)

    , n 是数组的大小。
	题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的。
	你可以按任意顺序返回答案。
*/

// TopKFrequent 前 K 个高频元素
func TopKFrequent(nums []int, k int) []int {
	ans:=[]int{}
	map_num:=map[int]int{}
	for _,item:=range nums {
		map_num[item]++
	}
	for key,_:=range map_num{
		ans=append(ans,key)
	}
	//核心思想：排序
	//可以不用包函数，自己实现快排
	sort.Slice(ans,func (a,b int)bool{
		return map_num[ans[a]]>map_num[ans[b]]
	})
	return ans[:k]
}

