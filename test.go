package main

import "fmt"

func main()  {
	n := 4
	chessboard := make([][]string, n)
	for i := 0; i < n; i++ {
		temp := make([]string,n)
		for j := 0;j < n;j++{
			temp[j] = "."
		}
		chessboard[i] = temp
	}

	fmt.Println(chessboard)
}
