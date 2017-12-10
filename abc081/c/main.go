package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

type BallType struct {
	Number string
	Count  int
}

type BallTypes []*BallType

func (b BallTypes) Len() int {
	return len(b)
}

func (b BallTypes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b BallTypes) Less(i, j int) bool {
	return b[i].Count < b[j].Count
}

func main() {
	sc.Scan()
	chunks := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(chunks[0])
	k, _ := strconv.Atoi(chunks[1])

	sc.Scan()
	a := strings.Split(sc.Text(), " ")

	cntByNum := map[string]int{}
	for _, b := range a {
		cntByNum[b]++
	}

	cnt := 0

	if n > k && len(cntByNum) > k {
		balls := BallTypes{}
		for n, c := range cntByNum {
			balls = append(balls, &BallType{Number: n, Count: c})
		}

		sort.Sort(sort.Reverse(balls))

		for i := k; i < len(balls); i++ {
			cnt += balls[i].Count
		}
	}

	fmt.Println(cnt)
}
