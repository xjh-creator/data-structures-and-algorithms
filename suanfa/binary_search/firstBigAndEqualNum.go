package main

// FirstBigAndEqualNumBinarySearch1 查找第一个值大于等于所给值的二分查找非递归实现
func FirstBigAndEqualNumBinarySearch1(array []int,n,value int) int {
	left,right := 0,n - 1

	for left <= right{
		mid := left + (right - left)/2
		if array[mid] >= value{
			if array[mid - 1] < value || mid == 0{
				return mid
			}
			right = mid - 1
		}else{
			left = mid + 1
		}
	}
	return -1
}

// bFirstBigAndEqualNumSearch 查找第一个值大于等于所给值的二分查找递归实现
func bFirstBigAndEqualNumSearch(array []int,n,value int) int {
	return FirstNumBinarySearch2(array,0,n - 1,value)
}

// FirstBigAndEqualNumBinarySearch2  查找第一个值大于等于所给值的二分查找递归实现
func FirstBigAndEqualNumBinarySearch2(array []int,left,right,value int) int {
	mid := left + (right - left)/2
	if array[mid] >= value{
		if array[mid - 1] < value || mid == 0{
			return mid
		}
		return FirstNumBinarySearch2(array,left,mid - 1,value)
	}else{
		return FirstNumBinarySearch2(array,mid + 1,right,value)
	}
	return -1
}

