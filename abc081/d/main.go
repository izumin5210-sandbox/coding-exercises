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
	sc.Scan()
	a := make([]int, n)
	for i, ai := range strings.Split(sc.Text(), " ") {
		a[i], _ = strconv.Atoi(ai)
	}

	steps := make([]string, 0, 2*n)

	first, last := 0, len(a)-1
	for last != first {
		sub := append([]int{}, a[first:last+1]...)
		sort.Sort(sort.Reverse(sort.IntSlice(sub)))

		if sub[1] > 0 || sub[0] > 0 && a[last] < sub[0] {
			max := math.MinInt64
			maxIdx := 0
			for i := first; i < last; i++ {
				if max < a[i] {
					max = a[i]
					maxIdx = i
				}
			}
			for max > a[last] {
				a[last] += max
				steps = append(steps, fmt.Sprintf("%d %d", maxIdx+1, last+1))
				// fmt.Printf("%d %d\n", maxIdx+1, last+1)
				// time.Sleep(1 * time.Second)
			}
			last--
		} else {
			min := math.MaxInt64
			minIdx := 0
			for i := first + 1; i <= last; i++ {
				if min > a[i] {
					min = a[i]
					minIdx = i
				}
			}
			for min < a[first] {
				a[first] += min
				steps = append(steps, fmt.Sprintf("%d %d", minIdx+1, first+1))
				// fmt.Printf("%d %d\n", minIdx+1, first+1)
				// time.Sleep(1 * time.Second)
			}
			first++
		}
	}

	fmt.Println(len(steps))
	fmt.Println(strings.Join(steps, "\n"))
}
