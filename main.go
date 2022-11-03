package main

import "fmt"

func main()  {
	var num int = 1
	temp := 'b' + num
	fmt.Println(temp)
	fmt.Println(string(temp))
	tempArray := []string{"a"}
	tempArray = append(tempArray,string(temp))
	fmt.Println(tempArray)
}
