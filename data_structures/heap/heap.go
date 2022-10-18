package main

import "fmt"

type heap struct {
	array []int //从小标1开始存储数据
	n     int   //堆中最多可以存储的元素个数
	count int   //堆中已经存储的数据个数
}

func newHeap(n int) *heap {
	return &heap{
		array: make([]int, 0, n+1),
		n:     n,
		count: 0,
	}
}

// insert 往堆添加元素，大顶堆
func (h *heap) insert(data int) {
	if h.count >= h.n {
		fmt.Println("堆满了")
		return
	}
	h.count++
	h.array[h.count] = data //先把数组插入到数组尾部，再自底向上堆化
	for i := h.count; i/2 > 0 && h.array[i] > h.array[i/2]; i /= 2 {
		h.array[i], h.array[i/2] = h.array[i/2], h.array[i]
	}
}

// removeMax 删除大顶堆的跟节点
func (h *heap) removeMax() {
	if h.count == 0 {
		fmt.Println("堆中没有元素")
		return
	}
	h.array[1], h.array[h.count] = h.array[h.count], 0
	h.count--
	heapify(&h.array, h.count, 1)
}

// heapify 自上向下堆化（先父节点，再左右子节点）, i 为自 i 下标的元素堆化
func heapify(array *[]int, count, i int) {
	for {
		maxPos := i
		if i*2 <= count && (*array)[i*2] > (*array)[i] {
			maxPos = i * 2
		}
		if i*2+1 <= count && (*array)[i*2+1] > (*array)[maxPos] {
			maxPos = i*2 + 1
		}
		if maxPos == i {
			break
		}
		(*array)[i], (*array)[maxPos] = (*array)[maxPos], (*array)[i]
		i = maxPos
	}
}

