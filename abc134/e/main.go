package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	res := make([]int, 0, n)

	min := math.MaxInt64
	max := -1

	for i := 0; i < n; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())

		var ok bool

		if v > min {
			for i, r := range res {
				if r >= v {
					continue
				}
				res[i] = v
				ok = true
				if v > max {
					max = v
				}
				break
			}
		}

		if !ok {
			res = append(res, v)
			if v < min {
				min = v
			}
		}
	}

	fmt.Println(len(res))
}
