package hash

import "sort"

/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

示例 1: 输入: s = "anagram", t = "nagaram" 输出: true

示例 2: 输入: s = "rat", t = "car" 输出: false

说明: 你可以假设字符串只包含小写字母。
*/

// IsAnagram 242.有效的字母异同位
func IsAnagram(s string,t string) bool{
	if len(s) != len(t){
		return false
	}
	record := [26]byte{}
	for _,v := range s{
		record[v - 'a']++
	}
	for _,v := range t{
		record[v - 'a']--
	}
	return record == [26]byte{}
}

/*
给你一个字符串数组 words ，请你找出所有在 words 的每个字符串中都出现的共用字符（包括重复字符），并以数组形式返回。你可以按 任意顺序 返回答案。

示例 1：

输入：words = ["bella","label","roller"] 输出：["e","l","l"] 示例 2：

输入：words = ["cool","lock","cook"] 输出：["c","o"]

提示：

1 <= words.length <= 100 1 <= words[i].length <= 100 words[i] 由小写英文字母组成
*/

// CommonChars 1002. 查找常用字符串
// 思路：每个字符串26位二进制（每位存字符串所含字符的出现次数），纵向比对，保留次数小的
func CommonChars(words []string) []string {
   length := len(words)
   tongji := [][26]int{}
	//统计词频
	for i:=0;i<length;i++{
		var row [26]int//存放该字符串的词频
		for j:=0;j<len(words[i]);j++{
			row[words[i][j]-97]++
		}
		tongji=append(tongji,row)
	}
	res:=make([]string,0)
	for i:=0;i<26;i++{ // 横向
		minTemp := tongji[0][i]
		for j:=1;j<len(tongji);j++{ // 纵向
			minTemp = min(minTemp,tongji[j][i])
		}
		if minTemp == 0{
			continue
		}
		for k:=0;k < minTemp;k++{
			str := string(rune(97 + i))
			res = append(res,str)
		}
	}
	return res
}

func min(a,b int) int {
	if a >= b{
		return b
	}
	return a
}

/*
	给定两个数组 nums1 和 nums2 ，
    返回 它们的交集。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。
*/

// intersection 349. 两个数组的交集
func intersection(nums1,nums2 []int) []int {
	set:=make(map[int]struct{},0)  // 用map模拟set
	res:=make([]int,0)
	for _,v:=range nums1{
		if _,ok:=set[v];!ok{
			set[v]=struct{}{}
		}
	}
	for _,v:=range nums2{
		//如果存在于上一个数组中，则加入结果集，并清空该set值
		if _,ok:=set[v];ok{
			res=append(res,v)
			delete(set, v)
		}
	}
	return res
}

/*
编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。如果 可以变为  1，那么这个数就是快乐数。

如果 n 是快乐数就返回 True ；不是，则返回 False 。

示例：

输入：19
输出：true
解释：
1^2 + 9^2 = 82
8^2 + 2^2 = 68
6^2 + 8^2 = 100
1^2 + 0^2 + 0^2 = 1
*/

// isHappy 202. 快乐数
func isHappy(n int) bool {
	getSum := func(num int) int{
		sum := 0
		for num > 0{
			sum += (num % 10) * (num % 10)
			num = num / 10
		}

		return sum
	}

	temp := make(map[int]bool,0)
	for n != 1 && !temp[n]{
		temp[n] = true
		n = getSum(n)
	}

	return n == 1
}

/*
	给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

	你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

	示例:

	给定 nums = [2, 7, 11, 15], target = 9

	因为 nums[0] + nums[1] = 2 + 7 = 9

	所以返回 [0, 1]
*/

// twoSum 1. 两数之和
func twoSum(nums []int, target int) []int {
	m := make(map[int]int,len(nums))

	for index, val := range nums {
		if preIndex, ok := m[target-val]; ok {
			return []int{preIndex, index}
		} else {
			m[val] = index
		}
	}

	return []int{}
}

/*
	给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

	注意： 答案中不可以包含重复的三元组。

	示例：

	给定数组 nums = [-1, 0, 1, 2, -1, -4]，

	满足要求的三元组集合为： [ [-1, 0, 1], [-1, -1, 2] ]
*/
// ThreeSum 三数之和
// 思路：
func ThreeSum(nums []int) [][]int {
	length := len(nums)
	if length < 3{
		return nil
	}

	// 1：排序数组
	sort.Ints(nums)

	result := make([][]int,0)
	for i:=0;i<length;i++{
		n1 := nums[i] // 2.第一层 for 循环第一个值 n1
		if n1 > 0{
			break
		}
		left,right := i+1,length-1
		for right > left{ // 3.第二层 for 循环 第二三个值 n2 n3
			n2,n3 := nums[left],nums[right]

			if n1 + n2 + n3 > 0{ // 4.数组是排序的，如果大于 0 ,就把最后一个最大的值往前挪一位
				right--
			}else if n1 + n2 + n3 < 0{ // 5.数组是排序的，如果小于于 0 ,就把 n2 值往后挪一位
				left++
			}else{
				result = append(result,[]int{nums[i],nums[left],nums[right]})
				for left < right && nums[left] == n2 { // 避免重复的值，如：n1 = 1,n2 + n3 = -1,
					left++                             // 如果往后一位 n2 还是同样的值，n3 也会是一样，造成重复
				}
				for left < right && nums[right] == n3 {
					right--
				}
			}
		}
	}
	return result
}
