package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	chunks := strings.Split(sc.Text(), " ")
	m, _ := strconv.Atoi(chunks[1])

	type city struct {
		pref int
		id   string
	}

	cities := []*city{}
	cityIdxByYear := make(map[int]int)
	years := []int{}

	for i := 0; i < m; i++ {
		sc.Scan()
		chunks := strings.Split(sc.Text(), " ")
		pref, _ := strconv.Atoi(chunks[0])
		year, _ := strconv.Atoi(chunks[1])
		cities = append(cities, &city{pref: pref})
		cityIdxByYear[year] = i
		years = append(years, year)
	}

	sort.IntSlice(years).Sort()
	idxByPref := make(map[int]int)

	for _, y := range years {
		city := cities[cityIdxByYear[y]]
		idxByPref[city.pref]++
		idx := idxByPref[city.pref]

		city.id = fmt.Sprintf("%06d%06d", city.pref, idx)
	}

	for _, c := range cities {
		fmt.Println(c.id)
	}
}
