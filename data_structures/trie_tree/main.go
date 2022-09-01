package main

import "fmt"

// SIZE 数组的大小
const SIZE = 26

// Trie 字典树
type Trie struct {
	*TrieNode
}

func NewTrie() *Trie {
	return &Trie{NewTrieNode('/')}
}

// TrieNode 字典树的节点
type TrieNode struct {
	data byte
	children []*TrieNode
	isEnding bool
}

func NewTrieNode(data byte) *TrieNode {
	return &TrieNode{
		data: data,
		children: make([]*TrieNode,SIZE),
	}
}

// insert 往字典树添加字符串
func (t *Trie)insert(text string)  {
	p := t.TrieNode
	for i := 0;i < len(text);i++ {
		index := text[i] - 'a'
		if p.children[index] == nil {
			newNode := NewTrieNode(text[i])
			p.children[index] = newNode
		}else{
			p.children[index].data = text[i]
		}
		p = p.children[index]
	}
	p.isEnding = true
}

// find 查找某个字符串是否存在于字典树中
func (t *Trie)find(pattern string) bool {
	if t.children == nil{
		return false
	}
	p := t.TrieNode
	for i := 0;i < len(pattern);i++ {
		index := pattern[i] - 'a'
		if p.children[index] == nil{
			return false
		}
		p = p.children[index]
	}
	if p.isEnding == false{
		return false
	}else {
		return true
	}
}

func main()  {
	root := NewTrie()
	fmt.Println(root.children[0])
	root.insert("abc")
	fmt.Println(root.find("ab"))
	fmt.Println(root.children[0].data)
	fmt.Println(root.children[0].children[1].data)
	fmt.Println(root.children[0].children[1].children[2].data)
}


