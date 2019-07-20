package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	dams := make([]int, n)

	for i := 0; i < n; i++ {
		sc.Scan()
		dams[i], _ = strconv.Atoi(sc.Text())
	}

	ms := make([]int, n)
	res := make([]string, n)

	for i := 0; i < n; i++ {
		a := 1
		for j := 0; j < n; j++ {
			ms[i] += a * dams[(i+j)%n]
			a *= -1
		}
		res[i] = fmt.Sprint(ms[i])
	}

	fmt.Println(strings.Join(res, " "))
}
