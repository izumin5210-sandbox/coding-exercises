package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	list := search(n)

	fmt.Println(len(list))
}

func search(n int) []int {
	d := int(math.Floor(math.Log10(float64(n)))) + 1
	if d < 4 {
		list := []int{}
		for _, p := range []int{357, 375, 537, 573, 735, 753} {
			if p <= n {
				list = append(list, p)
			}
		}
		return list
	}

	m, _ := strconv.Atoi(strings.Repeat("9", d-1))

	prev := search(m)
	list := make([]int, 0, 10*len(prev))

	min, _ := strconv.Atoi(strings.Repeat("9", d-2))

	for _, p := range prev {
		if p < min {
			continue
		}
		for _, q := range []int{3, 5, 7} {
			for i := 0; i < d; i++ {
				f := int(math.Pow10(i))
				r := f*q + p%f + (p/f)*10*f
				if r <= n {
					list = append(list, r)
				}
			}
		}
	}

	set := make(map[int]struct{}, len(list))
	for _, n := range list {
		set[n] = struct{}{}
	}
	newList := make([]int, 0, len(set))
	for p := range set {
		newList = append(newList, p)
	}

	return append(newList, prev...)
}
