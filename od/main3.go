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
	var solve func(s,t,minStep int)
	solve = func(s, t,minStep int) {
		if s == t{
			result = append(result,minStep)
			return
		}
		if s > t{
			return
		}
		solve(s + b,t,minStep)
		solve(s - b,t,minStep)
		minStep++
		solve(s - a,t,minStep)
		solve(s + a,t,minStep)
		minStep--
	}
	solve(s,t,0)
	if len(result) > 0{
		sort.Ints(result)
		fmt.Println(result[0])
	}
	fmt.Println(1)
}
