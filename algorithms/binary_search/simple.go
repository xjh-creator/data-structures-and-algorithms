package main

// mid := (left + right)/2
// mid := left + (right - left)/2
// mid := low + ((right - left)>>1)

// SimpleBinarySearch1 简单二分查找非递归实现
func SimpleBinarySearch1(array []int,n,value int) int {
	left,right := 0,n - 1

	for left <= right{
		mid := left + (right - left)/2
		if array[mid] == value{
			return mid
		}else if array[mid] < value{
			left = mid + 1
		}else{
			right = mid - 1
		}
	}
	return -1
}

// bSearch 简单二分查找递归实现
func bSimpleSearch(array []int,n,value int) int {
	return SimpleBinarySearch2(array,0,n - 1,value)
}

// SimpleBinarySearch2 简单二分查找递归实现
func SimpleBinarySearch2(array []int,left,right,value int) int {
	mid := left + (right - left)/2
	if array[mid] == value{
		return mid
	}else if array[mid] < value{
		return SimpleBinarySearch2(array,mid + 1,right,value)
	}else{
		return SimpleBinarySearch2(array,left,mid - 1,value)
	}
	return -1
}
