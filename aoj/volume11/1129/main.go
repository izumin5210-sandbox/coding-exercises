package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

func main() {
	for sc.Scan() {
		if sc.Text() == "0 0" {
			break
		}
		chunks := strings.Split(sc.Text(), " ")
		n, _ := strconv.Atoi(chunks[0])
		r, _ := strconv.Atoi(chunks[1])
		cards := make([]int, n, n)
		for i := 0; i < n; i++ {
			cards[i] = n - i
		}
		for i := 0; i < r; i++ {
			sc.Scan()
			chunks := strings.Split(sc.Text(), " ")
			p, _ := strconv.Atoi(chunks[0])
			c, _ := strconv.Atoi(chunks[1])
			block := make([]int, c, c)
			for j := 0; j < c; j++ {
				block[j] = cards[p+j-1]
			}
			for j := p - 2; j >= 0; j-- {
				cards[j+c] = cards[j]
			}
			for j := 0; j < c; j++ {
				cards[j] = block[j]
			}
		}
		fmt.Println(cards[0])
	}
}
