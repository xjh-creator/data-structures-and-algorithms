package main

// LastSmallAndEqualNumBinarySearch1 查找最后一个值小于等于所给值的二分查找非递归实现
func LastSmallAndEqualNumBinarySearch1(array []int,n,value int) int {
	left,right := 0,n - 1

	for left <= right{
		mid := left + (right - left)/2
		if array[mid] > value{
			right = mid - 1
		}else{
			if array[mid + 1] > value || mid == n - 1{
				return mid
			}
			left = mid + 1
		}
	}
	return -1
}

// bLastNumSearch 查找最后一个值小于等于所给值的二分查找递归实现
func bLastSmallAndEqualNumSearch(array []int,n,value int) int {
	return LastNumBinarySearch2(array,0,n - 1,value)
}

// LastSmallAndEqualNumBinarySearch2  查找最后一个值小于等于所给值的二分查找递归实现
func LastSmallAndEqualNumBinarySearch2(array []int,left,right,value int) int {
	mid := left + (right - left)/2
	if array[mid] > value{
		return LastNumBinarySearch2(array,left,mid - 1,value)
	}else{
		if array[mid + 1] > value || mid == right{
			return mid
		}
		return LastNumBinarySearch2(array,mid + 1,right,value)
	}
	return -1
}

