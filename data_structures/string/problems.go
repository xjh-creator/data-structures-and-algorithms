package string

import "fmt"

/*
	给定一个字符串 s 和一个整数 k，从字符串开头算起, 每计数至 2k 个字符，就反转这 2k 个字符中的前 k 个字符。

	如果剩余字符少于 k 个，则将剩余字符全部反转。

	如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
	示例:

	输入: s = "abcdefg", k = 2
	输出: "bacdfeg"
*/
func ReverseStr(s string, k int) string {
	ss := []byte(s)
	length := len(s)
	for i := 0; i < length; i += 2 * k {
		// 1. 每隔 2k 个字符的前 k 个字符进行反转
		// 2. 剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符
		if i + k <= length {
			reverse(ss[i:i+k])
		} else {
			reverse(ss[i:length])
		}
	}
	return string(ss)
}

func reverse(b []byte) {
	left := 0
	right := len(b) - 1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

func ReplaceSpace(s string) string {
	length := len(s)
	spaceNums := 0
	for i:=0;i<length;i++{
		if s[i] == ' '{
			spaceNums++
		}
	}
	tempSlice := make([]byte,0)
	for i:=0;i<length;i++{
		if s[i] == ' '{
			tempSlice = append(tempSlice,'%','2','0')
			continue
		}
		tempSlice = append(tempSlice,s[i])

	}
	return string(tempSlice)
}

/*
	给定一个字符串，逐个翻转字符串中的每个单词。

	示例 1：
	输入: "the sky is blue"
	输出: "blue is sky the"

	示例 2：
	输入: "  hello world!  "
	输出: "world! hello"
	解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。

	示例 3：
	输入: "a good   example"
	输出: "example good a"
	解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
*/
func ReverseWords(s string) string {
	temp := []byte(s)
	// 处理多余的空格
	fastIndex,lowIndex := 0,0
	for temp[fastIndex] == ' '{ // 最前
		fastIndex++
	}
	length := len(temp)
	for ;fastIndex<length;fastIndex++{ // 处理中间重复的空格
		if fastIndex-1 > 0 && temp[fastIndex-1] == temp[fastIndex] && temp[fastIndex] == ' ' {
			continue
		}
		temp[lowIndex] = temp[fastIndex]
		lowIndex++
	}
	//删除尾部冗余空格
	if lowIndex-1 > 0 && temp[lowIndex-1] == ' ' {
		temp = temp[:lowIndex-1]
	} else {
		temp = temp[:lowIndex]
	}
	s = ""
	// 从后遍历，把字符加到 s 中
	nums := 0
	for i := len(temp) - 1;i >= 0;i--{
		nums++
		if temp[i] == ' '{
			s = s + string(temp[i+1:i+nums]) + string(temp[i])
			nums = 0
		}
		if i == 0{
			s = s + string(temp[i:i+nums])
		}
	}
	return s
}

/*
字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。请定义一个函数实现字符串左旋转操作的功能。比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。

示例 1：
输入: s = "abcdefg", k = 2
输出:"cdefgab"

示例 2：
输入: s = "lrloseumgh", k = 6
输出:"umghlrlose"

限制：
1 <= k < s.length <= 10000
*/

func ReverseLeftWords(s string, n int) string {
	if len(s) < n{
		return ""
	}
	temp := []byte(s)
	return string(append(temp[n:],temp[:n]...))
}

/*
	给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。

示例 1:
输入: "abab"
输出: True
解释: 可由子字符串 "ab" 重复两次构成。

示例 2:
输入: "aba"
输出: False

示例 3:
输入: "abcabcabcabc"
输出: True
解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)
*/

func RepeatedSubstringPattern(s string) bool {
	length := len(s)
	if length < 2{
		return false
	}
	tempS := []byte(s)
	for i:=0;i<length;i++{
		childLen := i + 1
		if length % childLen > 0{
			continue
		}
		temp := ""
		nums := 0
		for{
			nums++
			temp = fmt.Sprintf("%s%s",temp,string(tempS[:i+1]))
			tempLen := len(temp)
			if tempLen == length{
				break
			}
		}
		fmt.Println(i,"---2",temp)
		if temp == s && nums > 1{
			return true
		}
		nums = 0
	}
	return false
}

/*
实现strStr()函数。

给定一个haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1: 输入: haystack = "hello", needle = "ll" 输出: 2

示例 2: 输入: haystack = "aaaaa", needle = "bba" 输出: -1

说明: 当needle是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。 对于本题而言，当needle是空字符串时我们应当返回 0 。这与C语言的strstr()以及 Java的indexOf( 定义相符。
*/

func StrStr(haystack string, needle string) int {
	lenHaystack := len(haystack)
	lenNeedle := len(needle)
	if lenHaystack < lenNeedle{
		return -1
	}

	for i:=0;i<lenHaystack;i++{
		isTrue := true
		if haystack[i] == needle[0]{
			temp := i
			j := 0
			for ;j<lenNeedle && temp < lenHaystack;j++{
				if haystack[temp] != needle[j]{
					isTrue = false
					break
				}
				temp++
			}
			if isTrue && j == lenNeedle{
				return i
			}
		}
	}
	return -1
}