package hash

/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

示例 1: 输入: s = "anagram", t = "nagaram" 输出: true

示例 2: 输入: s = "rat", t = "car" 输出: false

说明: 你可以假设字符串只包含小写字母。
*/

// IsAnagram 有效的字母异同位
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

// CommonChars 查找常用字符串
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
