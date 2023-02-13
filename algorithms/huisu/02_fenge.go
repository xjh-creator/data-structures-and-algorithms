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

func RestoreIpAddresses(s string) []string {
	res := make([]string,0)
    length := len(s)
	var dfs func(start int,path []string)
	dfs = func(start int,path []string) {
		if len(path) == 4{
			if start == length{
				temp := strings.Join(path,".")
				res = append(res,temp)
			}
			return
		}

		for i:=start;i<length;i++{
			if i != start && s[start] == '0' { // 含有前导 0，无效
				break
			}
			temp := s[start:i+1]
			if num,_ := strconv.Atoi(temp);num > 255 || num < 0{
				break
			}
			path = append(path,temp)
			dfs(i+1,path)
			path = path[:len(path)-1]
		}
	}
	path := make([]string,0)
	dfs(0,path)
	return res
}