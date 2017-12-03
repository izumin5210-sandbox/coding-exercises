package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	f := map[int]map[int]bool{}
	p := map[int]map[int]int{}
	cntByShop := map[int]int{}

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	for i := 0; i < n; i++ {
		f[i+1] = map[int]bool{}
		sc.Scan()
		row := sc.Text()
		for j, c := range strings.Split(row, " ") {
			f[i+1][j+1] = c == "1"
			if c == "1" {
				cntByShop[i+1]++
			}
		}
	}
	for i := 0; i < n; i++ {
		p[i+1] = map[int]int{}
		sc.Scan()
		row := sc.Text()
		for j, c := range strings.Split(row, " ") {
			if j+1 <= cntByShop[i+1] {
				break
			}
			v, _ := strconv.Atoi(c)
			p[i+1][j] = v
		}
	}

	for i, pi := range p {
		for j, v := range pi {
		}
	}
}
