package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

func main() {
	for sc.Scan() {
		n, _ := strconv.Atoi(sc.Text())
		if n == 0 {
			break
		}
		scores := make([]int, n, n)
		for i := 0; i < n; i++ {
			sc.Scan()
			scores[i], _ = strconv.Atoi(sc.Text())
		}
		sort.IntSlice(scores).Sort()
		var sum = 0
		for i := 1; i < len(scores)-1; i++ {
			sum += scores[i]
		}
		fmt.Println(sum / (n - 2))
	}
}
