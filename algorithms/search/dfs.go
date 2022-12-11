package search

import (
	"fmt"
	"sort"
)

/* 问题：假设我们都是计算机系的大学生，需要选修一些课程。但是要选修有的课程必须要先学习它的前序课程。例如，学习网络首先要学习操作系统的知识，
   而要学习操作系统的知识必须首先学习数据结构的知识。
   如果我们现在只知道每门课程的前序课程，不清楚完整的学习路径，
   我们要怎么设计这一系列课程学习的顺序，确保我们在学习任意一门课程的时候，
   都已 学完了它的前序课程呢？
*/

// 计算机课程和其前序课程的映射关系
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus": {"linear algebra"},
	"compilers": { "data structures", "formal languages", "computer organization", },
	"data structures": {"discrete math"},
	"databases": {"data structures"},
	"discrete math": {"intro to programming"},
	"formal languages": {"discrete math"},
	"networks": {"operating systems"},
	"operating systems": {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}



func DFS() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	// 匿名函数，用以查找并整合一整条课程学习顺序路线
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item]) // 这里类似于压栈，先传入的是 prereqs 中非前缀的部分，然后深入遍历，一直找到最先需要学的课程
				order = append(order, item) // 递归到最后，栈最上面的就是最先需要学的课程，最新把它加入到返回结果这个数组中
			}
		}
	}
	var keys []string
	// 拿到 map 的 key 值，也就是计算机课程（非前缀课程）
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

