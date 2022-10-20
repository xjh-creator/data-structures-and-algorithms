package main

import (
	"errors"
	"fmt"
)

//稀疏图 - 邻接表
type sparseGraph struct {
	nodeNum  int  //节点数
	edgeNum  int  //边数量
	directed bool //有向 or 无向图
	graph    [][]int
}

func NewSparseGraph(nodeNum int, directed bool) *sparseGraph {
	buf := make([][]int, nodeNum)
	return &sparseGraph{
		nodeNum:  nodeNum,
		edgeNum:  0,
		directed: directed,
		graph:    buf,
	}
}

// GetVertex 获取顶点数量
func (s *sparseGraph) GetVertex() int {
	return s.nodeNum
}

// GetEdge 获取边数量
func (s *sparseGraph) GetEdge() int {
	return s.edgeNum
}

// AddEdge 添加边: v1,v2表示顶点相应的索引
func (s *sparseGraph) AddEdge(v1, v2 int) error {
	// 判断索引是否越界
	if v1 < 0 || v2 < 0 || v1 >= s.nodeNum || v2 >= s.nodeNum {
		return errors.New("index is illegal.")
	}

	// 没有处理平行边的情况
	s.graph[v1] = append(s.graph[v1], v2)
	if v1 != v2 && !s.directed {
		s.graph[v2] = append(s.graph[v2], v1)
	}
	s.edgeNum++
	return nil
}

// HasEdge 判断v1,v2是否已经有边
func (s *sparseGraph) HasEdge(v1, v2 int) (bool, error) {
	// 判断索引是否越界
	if v1 < 0 || v2 < 0 || v1 >= s.nodeNum || v2 >= s.nodeNum {
		return false, errors.New("index is illegal.")
	}

	for i := 0; i < len(s.graph[v1]); i++ {
		if s.graph[v1][i] == v2 {
			return true, nil
		}
	}
	return false, nil
}

// 迭代器: 输出节点v所连接的节点,时间复杂度为O(1)
func (s *sparseGraph) Iterator(v int) []int {
	// 判断索引是否越界
	if v < 0 || v >= s.nodeNum {
		fmt.Println("index is illegal.")
		return nil
	}
	return s.graph[v]
}
