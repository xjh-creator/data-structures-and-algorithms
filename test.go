package main

func reverseWords(s string) string {
	// 1.去除多余的空格
	// 2.把整个字符串倒转
	// 3. 倒转单词

	//1.使用双指针删除冗余的空格
	slowIndex, fastIndex := 0, 0
	temp := []byte(s)
	//删除头部冗余空格
	for len(temp) > 0 && fastIndex < len(temp) && temp[fastIndex] == ' ' {
		fastIndex++
	}
	//删除单词间冗余空格
	for ; fastIndex < len(temp); fastIndex++ {
		if fastIndex-1 > 0 && temp[fastIndex-1] == temp[fastIndex] && temp[fastIndex] == ' ' {
			continue
		}
		temp[slowIndex] = temp[fastIndex]
		slowIndex++
	}
	//删除尾部冗余空格
	if slowIndex-1 > 0 && temp[slowIndex-1] == ' ' {
		temp = temp[:slowIndex-1]
	} else {
		temp = temp[:slowIndex]
	}

	// 2.反转整个字符串
	reverse(temp,0,len(temp) - 1)

	// 3.反转单词
	start := 0
	for i:=0;i<len(temp);i++{
		if temp[i] == ' '{
			reverse(temp,start,i-1)
			start = i + 1
		}
		if i == len(temp) - 1{
			reverse(temp,start,i)
		}
	}

	return string(temp)
}

func reverse(b []byte,left,right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

func main()  {


}
