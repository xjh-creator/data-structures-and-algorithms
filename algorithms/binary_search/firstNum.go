package main

// FirstNumBinarySearch1 查找第一个值等于所给值的二分查找非递归实现
func FirstNumBinarySearch1(array []int,n,value int) int {
	left,right := 0,n - 1

	for left <= right{
		mid := left + (right - left)/2
		if array[mid] == value{
			if array[mid - 1] != value || mid == 0{
				return mid
			}
			right = mid - 1
		}else if array[mid] < value{
			left = mid + 1
		}else{
			right = mid - 1
		}
	}
	return -1
}

// bFirstNumSearch 查找第一个值等于所给值的二分查找递归实现
func bFirstNumSearch(array []int,n,value int) int {
	return FirstNumBinarySearch2(array,0,n - 1,value)
}

// FirstNumBinarySearch2  查找第一个值等于所给值的二分查找递归实现
func FirstNumBinarySearch2(array []int,left,right,value int) int {
	mid := left + (right - left)/2
	if array[mid] == value{
		if array[mid - 1] != value || mid == 0{
			return mid
		}
		return FirstNumBinarySearch2(array,left,mid - 1,value)
	}else if array[mid] < value{
		return FirstNumBinarySearch2(array,mid + 1,right,value)
	}else{
		return FirstNumBinarySearch2(array,left,mid - 1,value)
	}
	return -1
}

