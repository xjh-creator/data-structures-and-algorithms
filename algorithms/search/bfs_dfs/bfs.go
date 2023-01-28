package bfs_dfs

import (
	"container/list"
	"fmt"
)

// Graph 图，用邻接表来实现
type Graph struct {
	adj []*list.List
	v int // 节点个数
}

func NewGraph(v int) *Graph {
	graph := Graph{
		v:v,
		adj:make([]*list.List,v),
	}
	for k := range graph.adj{
		graph.adj[k] = new(list.List)
	}
	return &graph
}

func (g *Graph)AddEdge(start,target int)  {
	g.adj[start].PushBack(target)
	g.adj[target].PushBack(start)
}

// BFS 宽度优先搜索算法，从一个起点，走到终点，问最短路径。
// start 起点 ， target 终点
func (graph *Graph)BFS(start,target int){
	if start == target{
		return
	}

	// hasAcross 已走过的路径
	hasAcross := make([]int,graph.v)
	for i := range hasAcross{
		hasAcross[i] = -1
	}

	// queue 邻接表的节点
	queue := make([]int,0)
	// isVisited 已经走过的节点
	isVisited := make([]bool,graph.v)
	queue = append(queue,start)
	isVisited[start] = true
	isFound := false
	for len(queue) > 0 && !isFound{
		top := queue[0]
		queue = queue[1:]
		linkList := graph.adj[top]
		for e := linkList.Front();e.Next() != nil;e = e.Next(){
			k := e.Value.(int)
			if !isVisited[k]{
				hasAcross[k] = top
				if k == target{
					isFound = true
					break
				}
				queue = append(queue,k)
				isVisited[k] = true
			}
		}
	}
	if isFound{
		printPrev(hasAcross,start,target)
	}else{
		fmt.Printf("no path found from %d to %d\n",start,target)
	}
}


//print path recursively
func printPrev(prev []int, s int, t int) {

	if t == s || prev[t] == -1 {
		fmt.Printf("%d ", t)
	} else {
		printPrev(prev, s, prev[t])
		fmt.Printf("%d ", t)
	}

}
