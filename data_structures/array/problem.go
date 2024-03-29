package array

import "sort"

//给定一个n个元素有序的（升序）整型数组 nums 和一个目标值 target  ，
//写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

// Search 二分查找 704
func Search(nums []int,target int) int {
	length := len(nums)
	left,right := 0,length - 1

	for left <= right{
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
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 O(log n) 的算法。

*/

// SearchInsert 搜索插入的位置 35
func SearchInsert(nums []int, target int) int {

	left,right := 0,len(nums) - 1

	ret := 0
	for left <= right{
		middle := left + (right - left) / 2
		if nums[middle] == target{
			return middle
		}else if nums[middle] > target{
			right = middle - 1
			ret = middle
		}else{
			left = middle + 1
			ret = middle + 1
		}
	}

	return ret
}

/*
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。

*/

// SearchRange 34. 在排序数组中查找元素的第一个和最后一个位置
func SearchRange(nums []int, target int) []int {
	left,right := 0,len(nums) - 1

	index := 0
	isFind := false
	for left <= right{
		middle := left + (right - left) / 2
		if nums[middle] == target{
			index = middle
			isFind = true
			break
		}else if nums[middle] > target{
			right = middle - 1
		}else{
			left = middle + 1
		}
	}
	if !isFind{
		return []int{-1,-1}
	}
	if index == 0{
		for right - 1 >= index && nums[index + 1] == target{
			index++
		}
		return []int{0,index}
	}
	right = index
	for len(nums) > right + 1 && nums[right + 1] == target{
		right++
	}
	left = index
	for left - 1 >= 0 && nums[left - 1] == target{
		left--
	}
	return []int{left,right}
}

/*
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于val的元素，并返回移除后数组的新长度。

不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并原地修改输入数组。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

示例1: 给定 nums = [3,2,2,3], val = 3, 函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。 你不需要考虑数组中超出新长度后面的元素。

示例2: 给定 nums = [0,1,2,2,3,0,4,2], val = 2, 函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。
*/

// RemoveElement 移除元素 27
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
给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。

由于在某些语言中不能改变数组的长度，所以必须将结果放在数组nums的第一部分。更规范地说，如果在删除重复项之后有 k 个元素，那么 nums 的前 k 个元素应该保存最终结果。

将最终结果插入 nums 的前 k 个位置后返回 k 。

不要使用额外的空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

输入：nums = [1,1,2]
输出：2, nums = [1,2,_]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
*/


// removeDuplicates 26 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

/*
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

示例 1： 输入：nums = [-4,-1,0,3,10] 输出：[0,1,9,16,100] 解释：平方后，数组变为 [16,1,0,9,100]，排序后，数组变为 [0,1,9,16,100]

示例 2： 输入：nums = [-7,-3,2,3,11] 输出：[4,9,9,49,121]
*/

// SortedSquares 977 有序数组的平方
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

// MinSubArrayLen 209.长度最小子数组
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

// generateMatrix 59.螺旋矩阵II
func generateMatrix(n int) [][]int {
	ret := make([][]int,n)
	for k := range ret{
		ret[k] = make([]int,n)
	}

	loop := n / 2 // loop 循环的次数
	count := 1 // 每一个空格的值，初始为 1
	offset := 1 // 左闭右开，右闭限制的位数
	startx,starty := 0,0

	for loop > 0{
		i,j := startx,starty

		// 从左往右,行数不变，列数在变
		for j = starty;j < n - offset;j++{
			ret[i][j] = count
			count++
		}

		// 从右侧从上往下,行数在变，列数不变
		for i = startx;i < n - offset;i++{
			ret[i][j] = count
			count++
		}

		// 从右往左，行数不变，列数在变
		for ;j > starty;j--{
			ret[i][j] = count
			count++
		}

		// 从下往上，行数在变，列数不变
		for ;i > startx;i--{
			ret[i][j] = count
			count++
		}
		startx++
		starty++
		offset++
		loop--
	}

	center := n / 2
	if n % 2 == 1{
		ret[center][center] = n * n
	}

	return ret
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

