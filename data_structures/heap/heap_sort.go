package main

// 堆排序分为建序跟排序
// 从上往下堆化：先父节点，再左右子节点

// buildHeap 建堆 , 从 count / 2 非叶子节点开始（因为从上往下堆化，叶子节点无从比较，所以从 count / 2 的非叶子节点开始）
// 建序后形成的堆，按照数组下标顺序排序尚不成序
func buildHeap(array *[]int, count int) {
	for i := count / 2; i >= 1; i-- {
		heapify(array, count, i)
	}
}

// sort 排序，count表示元素的个数，数组的元素从下标1到n的过程
func sort(array *[]int, count int) {
	buildHeap(array, count)
	k := count
	for k > 1 {
		(*array)[1], (*array)[k] = (*array)[k], (*array)[1]
		k--
		heapify(array, k, 1)
	}
}
