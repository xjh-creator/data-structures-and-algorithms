package sort

func MergeSort(array []int)  {
	MergeSortDetail(array,0,len(array) - 1)
}

func MergeSortDetail(array []int,l,r int)  {
	if l >= r{
		return
	}
	middle := (l + r) /2
	MergeSortDetail(array,l,middle)
	MergeSortDetail(array,middle + 1,r)
	if array[middle] > array[middle +1]{
		Merge(array,l,middle,r)
	}
}

func Merge(array []int,l,mid,r int)  {
	temp := make([]int,r + 1)
	for i:=l;i<=r;i++{
		temp[i] = array[i]
	}

	i,j := l,mid + 1
	for k:=l;k<=r;k++{
		if i > mid{
			array[k] = temp[j]
			j++
		}else if j > r{
			array[k] = temp[i]
			i++
		}else if temp[i] > temp[j]{
			array[k] = temp[j]
			j++
		}else{
			array[k] = temp[i]
			i++
		}
	}
}