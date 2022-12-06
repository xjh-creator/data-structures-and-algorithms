package sort

func BubbleSort(array []int)  {
	n := len(array)
	if n <= 1{
		return
	}

	for i := 0;i < n ;i++{
        flag := false
		for j := 1;j < n - i;j++{
			if array[j] < array[j - 1]{
				array[j],array[j-1] = array[j-1],array[j]
				flag = true
			}
		}
		if !flag{
			break
		}
	}
}
