package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main()  {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str := strings.Split(input.Text()," ")

	s,_ := strconv.Atoi(str[0])
	t,_ := strconv.Atoi(str[1])
	a,_ := strconv.Atoi(str[2])
	b,_ := strconv.Atoi(str[3])
	if s == t{
		fmt.Println(0)
	}

	result := []int{}
	minStep := 0
	var solve func(s,t int)
	solve = func(s, t int) {
		if s == t{
			result = append(result,minStep)
			return
		}
		if s > t{
			return
		}
		if s + b < t{
			solve(s + b,t)
		}
		if s + a < t{
			minStep++
			solve(s + a,t)
			minStep--
		}
		if s - b > t{
			solve(s - b,t)
		}
		minStep++
		solve(s - a,t)
		minStep--
	}
	solve(s,t)
	if len(result) > 0{
		sort.Ints(result)
		fmt.Println(result)
	}
	fmt.Println(1)
}
