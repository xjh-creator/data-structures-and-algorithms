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
