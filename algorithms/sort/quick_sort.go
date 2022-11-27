package sort

func QuickSort(array []int) {
	sortDetail(array, 0, len(array)-1)
}

// sortDetail l , r 分别为数组的最左、最右下标
func sortDetail(array []int, l, r int) {
	//递归一定要有截止条件
	if l >= r {
		return
	}
	point := partition(array, l, r)
	sortDetail(array, l, point-1)
	sortDetail(array, point+1, r)
}

// partition 分区函数
func partition(array []int, l, r int) int {
	//pivot 为分区点，这里选择切片的最后一个元素作为分区点
	point := array[r]
	//i 元素为数组最左元素的下标
	i := l
	for j := l; j < r; j++ {
		if array[j] < point{
			if !(i == j){
				array[i], array[j] = array[j], array[i]
			}
			i++
		}
	}
	//上面 for 循环只是到 r - 1,这里处理最后一个元素 pivot
	array[i], array[r] = array[r], array[i]
	return i
}
