package main

import (
	"errors"
	"fmt"
)

// 邻接矩阵实现无权图
// 稠密图 - 邻接矩阵
type denseGraph struct {
	n        int  //节点数
	m        int  //边数
	directed bool //有向图 or 无向图
	graph    [][]bool
}

// NewDenseGraph 构造函数:有n个顶点,有向 or 无向图
func NewDenseGraph(n int, directed bool) *denseGraph {
	// 初始化 n*n 的二维切片矩阵
	buf := make([][]bool, n)
	for i := 0; i < n; i++ {
		buf[i] = make([]bool, n)
	}
	return &denseGraph{
		n:        n,
		m:        0,
		directed: directed,
		graph:    buf,
	}
}

// GetVertex 获取顶点数量
func (d denseGraph) GetVertex() int {
	return d.n
}

// GetEdge 获取边数量
func (d denseGraph) GetEdge() int {
	return d.m
}

// AddEdge 添加边: v1,v2均表示顶点相应的索引
func (d *denseGraph) AddEdge(v1, v2 int) error {
	b, err := d.HasEdge(v1, v2)
	if err != nil {
		return err
	}
	// 不要平行边
	if !b {
		// 表示v1 -> v2有条边
		d.graph[v1][v2] = true
		if !d.directed {
			// 如果是无向图,v2 -> v1也要表示有边
			d.graph[v2][v1] = true
		}
		d.m++
	}
	return nil
}

// HasEdge 判断v1,v2是否已经有边
func (d *denseGraph) HasEdge(v1, v2 int) (bool, error) {
	// 判断索引是否越界
	if v1 < 0 || v2 < 0 || v1 >= d.n || v2 >= d.n {
		return false, errors.New("index is illegal.")
	}
	return d.graph[v1][v2], nil
}

// Iterator 迭代器: 输出节点v所连接的节点,时间复杂度为O(n)
func (d *denseGraph) Iterator(v int) []int {
	// 判断索引是否越界
	if v < 0 || v >= d.n {
		fmt.Println("index is illegal.")
		return nil
	}

	buf := []int{}
	for i, j := range d.graph[v] {
		if j {
			buf = append(buf, i)
		}
	}
	return buf
}
