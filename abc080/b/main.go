package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Scan()
	str := sc.Text()
	n, _ := strconv.Atoi(str)

	sum := 0
	for _, c := range str {
		n, _ := strconv.Atoi(string(c))
		sum += n
	}

	if n%sum == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
