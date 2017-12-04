package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

func main() {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	for i := 0; i < n; i++ {
		sc.Scan()
		edges := make([]int, 3, 3)
		for i, c := range strings.Split(sc.Text(), " ") {
			edges[i], _ = strconv.Atoi(c)
			edges[i] = int(math.Pow(float64(edges[i]), 2))
		}
		sort.IntSlice(edges).Sort()
		msg := "NO"
		if edges[2] == edges[0]+edges[1] {
			msg = "YES"
		}
		fmt.Println(msg)
	}
}
