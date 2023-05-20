package main

import (
	"strconv"
	"strings"
)

/*
	给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。

	返回 s 所有可能的分割方案。回文串 是正着读和反着读都一样的字符串。

	示例: 输入: "aab" 输出: [ ["aa","b"], ["a","a","b"] ]
    示例： 输入：s = "a" 	输出：[["a"]]
*/
// Partition 131.分割回文串
func Partition(s string) [][]string {
	res := make([][]string,0)
	path := make([]string,0)
    length := len(s)
	var dfs func(start int)
	dfs = func(start int) {
		if start == length{
			temp := make([]string,len(path))
			copy(temp,path)
			res = append(res,temp)
			return
		}

		for i:=start;i<length;i++{
			temp := s[start:i+1]
			if judge(temp){
				path = append(path,temp)
				dfs(i+1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return res
}

func judge(s string) bool {
	if s == ""{
		return false
	}
	length := len(s)
	if length == 1{
		return true
	}
	left,right := 0,length - 1
	for left < right{
		if s[left] != s[right]{
			return false
		}
		left++
		right--
	}
	return true
}

/*
给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

有效的 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。

例如："0.1.2.201" 和 "192.168.1.1" 是 有效的 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效的 IP 地址。

示例 1：

    输入：s = "25525511135"
    输出：["255.255.11.135","255.255.111.35"]

示例 2：

    输入：s = "0000"
    输出：["0.0.0.0"]

*/

// RestoreIpAddresses 93.复原IP地址
func RestoreIpAddresses(s string) []string {
	res,path := make([]string,0),make([]string,0,len(s))

	var dfs func(start int)
	dfs = func(start int){
		if len(path) == 4{
			if start == len(s){
				temp := strings.Join(path,".")
				res = append(res,temp)
			}

			return
		}

		for i:=start;i<len(s);i++{
			if i != start && s[start] == '0'{
				break
			}

			str := s[start:i+1]
			num,_ := strconv.Atoi(str)
			if num > 255{
				break
			}
			path = append(path,str)
			dfs(i + 1)
			path = path[:len(path) - 1]
		}
	}

	dfs(0)

	return res
}