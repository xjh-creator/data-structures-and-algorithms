package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str1 := input.Text()

	input.Scan()
	str2 := input.Text()

	len1 := len(str1)
	len2 := len(str2)
	if len1 < len2{
		str1,str2 = str2,str1
		len1,len2 = len2,len1
	}

	maxChildStr := []byte{}
    lenMax := 0
    for i := range str1{
    	for j := range str2{
    		if str1[i] == str2[j]{
    			tempI := i
    			tempJ := j
    			tempChild := []byte{}
    			for str1[tempI] == str2[tempJ]{
					tempChild = append(tempChild,str2[tempJ])
					if tempI + 1 >= len1 || tempJ + 1 >= len2{
						break
					}
					tempI++
    				tempJ++
				}
				tempLen := len(tempChild)
				if tempLen > lenMax{
					maxChildStr = tempChild
					lenMax = tempLen
				}
			}
		}
	}
	if len(maxChildStr) > 0{
		fmt.Println(string(maxChildStr))
	}else{
		fmt.Println("")
	}

}
