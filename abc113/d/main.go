package main

import "fmt"

func main() {
	var h, w, k int
	fmt.Scan(&h, &w, &k)

	cnt := search(h, w, k, 1)

	fmt.Println(cnt)
}

func search(h, w, k, s int) int {
	if h == 0 {
		if s == k {
			return 1
		}
		return 0
	}

	var c int
	if s > 1 {
		c += count(s-2) * search(h-1, w, k, s-1)
	}
	c += search(h-1, w, k, s)
	if s < w {
		c += count(w-s-1) * search(h-1, w, k, s+1)
	}
	return c
}

func count(in int) int {
	if in > 1 {
		return count(in-2) + count(in-1)
	}
	if in == 1 {
		return 2
	}
	return 1
}
