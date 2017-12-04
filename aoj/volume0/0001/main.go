package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	n  = 10
	sc = bufio.NewScanner(os.Stdin)
)

func main() {
	heights := make([]int, n, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		heights[i], _ = strconv.Atoi(sc.Text())
	}
	sort.Sort(sort.Reverse(sort.IntSlice(heights)))

	for i := 0; i < 3; i++ {
		fmt.Println(heights[i])
	}
}
