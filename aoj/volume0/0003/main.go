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
		chunks := strings.Split(sc.Text(), " ")
		a, _ := strconv.Atoi(chunks[0])
		b, _ := strconv.Atoi(chunks[1])
		fmt.Println(len(fmt.Sprint(a + b)))
	}
}
