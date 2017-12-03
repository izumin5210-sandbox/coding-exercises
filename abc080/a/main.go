package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Scan()
	chunks := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(chunks[0])
	a, _ := strconv.Atoi(chunks[1])
	b, _ := strconv.Atoi(chunks[2])

	fmt.Println(int(math.Min(float64(n*a), float64(b))))
}
