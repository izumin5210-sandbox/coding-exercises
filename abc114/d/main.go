package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := n - 1; i > 0; i-- {
		n = n * i
	}

	var list []int
	for i := 1; i < n/2+1; i++ {
		if n%i == 0 {
			list = append(list, i)
		}
	}

	for _, m := range list {
		if countDevisor(m) == 75 {
			fmt.Println(m)
		}
	}
}

func countDevisor(n int) int {
	if c, ok := memo[n]; ok {
		return c
	}

	var c int

	for i := 1; i < n/2+1; i++ {
		if n%i == 0 {
			c++
		}
	}

	memo[n] = c
	return c
}

var memo map[int]int

func init() {
	memo = make(map[int]int)
}
