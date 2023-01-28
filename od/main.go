package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	firstInput := input.Text()
	firstInput = strings.TrimPrefix(firstInput,"[")
	firstInput = strings.TrimSuffix(firstInput,"]")
	// array 输入的数组
	array := strings.Split(firstInput,",")
	input.Scan()
	// nums 输入的数量
	nums,_ := strconv.Atoi(input.Text())
	if !(nums == 1 || nums == 2 || nums == 4 || nums == 8){
		return
	}

	firstList := make([]string,0)
	secondList := make([]string,0)
	for _,v := range array{
		if v == "0" || v == "1" || v == "2" || v == "3"{
			firstList = append(firstList,v)
		}else{
			secondList = append(secondList,v)
		}
	}
	fLen := len(firstList)
	sLen := len(secondList)

	result := make([][]string,0)
	if nums == 8{
		result = append(result,firstList,secondList)
	}else if nums == 4{
		if fLen == 4{
			result = append(result,firstList)
		}else if sLen == 4{
			result = append(result,secondList)
		}
	}else if nums == 2{
		if fLen == 2{
			result = append(result,firstList)
		}else if sLen == 2{
			result = append(result,secondList)
		}else if fLen == 4{
			result = getFourResult(firstList)
		}else if sLen == 4{
			result = getFourResult(secondList)
		}else if fLen == 3{
			result = getTreeResult(firstList)
		}else if sLen == 3{
			result = getTreeResult(secondList)
		}
	}else{
		if fLen == 1{
			result = append(result,firstList)
		}else if sLen == 1{
			result = append(result,secondList)
		}else if fLen == 3{
			result = [][]string{
				{firstList[0]},
				{firstList[1]},
				{firstList[2]},
			}
		}else if sLen == 3{
			result = [][]string{
				{secondList[0]},
				{secondList[1]},
				{secondList[2]},
			}
		}else if fLen == 2{
			result = [][]string{
				{firstList[0]},
				{firstList[1]},
			}
		}else if sLen == 2{
			result = [][]string{
				{secondList[0]},
				{secondList[1]},
			}
		}else if fLen == 4{
			result = [][]string{
				{firstList[0]},
				{firstList[1]},
				{firstList[2]},
				{firstList[3]},
			}
		}else if sLen == 4{
			result = [][]string{
				{secondList[0]},
				{secondList[1]},
				{secondList[2]},
				{secondList[3]},
			}
		}
	}
	fmt.Println(result)
}



func getTreeResult(array []string) [][]string {
	return [][]string{
		{array[0],array[1]},{array[0],array[2]},
		{array[1],array[2]},
	}
}

func getFourResult(array []string) [][]string {
	return [][]string{
		{array[0],array[1]},{array[0],array[2]},{array[0],array[3]},
		{array[1],array[2]},{array[1],array[3]},
		{array[2],array[3]},
	}
}

