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
		p, _ := strconv.Atoi(chunks[1])
		cnts := make([]int, n, n)
		for i := 0; true; i++ {
			j := i % n
			if p == 0 {
				p, cnts[j] = cnts[j], 0
			} else {
				cnts[j]++
				p--
			}

			if p == 0 {
				noneCnt := 0
				winner := 0
				for j, c := range cnts {
					if c == 0 {
						noneCnt++
					} else {
						winner = j
					}
				}
				if noneCnt == n-1 {
					fmt.Println(winner)
					break
				}
			}
		}
	}
}
