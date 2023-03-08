package double_pointer

import "sort"

/*
	给你一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，并返回移除后数组的新长度。

	不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并原地修改输入数组。

	元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
*/
// removeElement 27. 移除元素
func removeElement(nums []int, val int) int {
	length := len(nums)

	res := 0
	for i:=0;i<length;i++{
		if nums[i] != val{
			nums[res] = nums[i]
			res++
		}
	}
	nums = nums[:res]
	return res
}

// removeDuplicates 26. 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	lowIndex := 0
	length := len(nums)
	for fastIndex:=0;fastIndex<length;fastIndex++{
		if nums[lowIndex] != nums[fastIndex]{
			lowIndex = lowIndex + 1
			nums[lowIndex] = nums[fastIndex]
		}

	}
	return lowIndex + 1
}

/*
	给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

	请注意 ，必须在不复制数组的情况下原地对数组进行操作。
*/
// moveZeroes 283. 移动零
func moveZeroes(nums []int)  {
	length := len(nums)
	if length == 1{
		return
	}

	low := 0
	for i:=0;i<length;i++{
		if nums[i] != 0{
			nums[low],nums[i] = nums[i],nums[low]
			low++
		}
	}
	return
}

/*
	给定 s 和 t 两个字符串，当它们分别被输入到空白的文本编辑器后，如果两者相等，返回 true 。# 代表退格字符。

    注意：如果对空文本输入退格字符，文本继续为空。

*/
// backspaceCompare 844. 比较含退格的字符串
func backspaceCompare(s string, t string) bool {
	sNums,tNums := 0,0
	i,j := len(s) - 1,len(t) - 1

	for i>=0 || j>=0{
		for i>=0{
			if s[i] == '#'{
				sNums++
				i--
			}else if sNums > 0{
				sNums--
				i--
			}else{
				break
			}
		}
		for j>=0{
			if s[j] == '#'{
				tNums++
				j--
			}else if tNums > 0{
				tNums--
				j--
			}else{
				break
			}
		}
		if i>=0 && j>=0{
			if s[i] != t[j]{
				return false
			}
		}else if i>=0 || j>=0{
			return false
		}
		i--
		j--
	}
	return true
}

// sortedSquares 977.有序数组的平方
func sortedSquares(nums []int) []int {

}


