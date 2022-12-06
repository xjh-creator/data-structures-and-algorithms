package sort

func InsertSort(array []int)  {
	n := len(array)
	if n <= 1{
		return
	}

	for i:=1;i<n;i++{
		value := array[i]
		j := i-1
		for ;j>0;j--{
			if array[j] > value{ // 移动元素
				array[j + 1] = array[j]
			}else{
				break
			}
		}
		array[j+1] = value // 插入数据
	}
}
