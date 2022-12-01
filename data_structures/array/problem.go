package array

//给定一个n个元素有序的（升序）整型数组 nums 和一个目标值 target  ，
//写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

// Search 二分查找
func Search(nums []int,target int) int {
	length := len(nums)
	left,right := 0,length - 1

	for left < right{
		middle := left + (right - left) / 2
		if nums[middle] == target{
			return middle
		}else if nums[middle] > target{
			right = middle - 1
		}else{
			left = middle + 1
		}
	}
	return -1
}
