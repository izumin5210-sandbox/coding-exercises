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

type Point struct {
	X, Y int
}

type Gem struct {
	ID int
	Point
}

type Map struct {
	Current       *Point
	Gems          map[int]*Gem
	CollectedGems map[int]*Gem
}

func (m *Map) Move(dir string, l int) {
	switch {
	case dir == "N" || dir == "S":
		a := 1
		if dir == "S" {
			a = -1
		}
		for i := 0; i < l; i++ {
			m.Current.Y += a
			for _, g := range m.Gems {
				m.TryToTakeGem(g)
			}
		}
	case dir == "E" || dir == "W":
		a := 1
		if dir == "W" {
			a = -1
		}
		for i := 0; i < l; i++ {
			m.Current.X += a
			for _, g := range m.Gems {
				m.TryToTakeGem(g)
			}
		}
	}
}

func (m *Map) HasCompleted() bool {
	return len(m.Gems) == 0
}

func (m *Map) TryToTakeGem(g *Gem) {
	if m.Current.X == g.X && m.Current.Y == g.Y {
		m.CollectedGems[g.ID] = g
		delete(m.Gems, g.ID)
	}
}

func main() {
	for sc.Scan() {
		n, _ := strconv.Atoi(sc.Text())
		if n == 0 {
			break
		}
		m := &Map{
			Current:       &Point{X: 10, Y: 10},
			Gems:          map[int]*Gem{},
			CollectedGems: map[int]*Gem{},
		}
		for i := 0; i < n; i++ {
			sc.Scan()
			chunks := strings.Split(sc.Text(), " ")
			x, _ := strconv.Atoi(chunks[0])
			y, _ := strconv.Atoi(chunks[1])
			m.Gems[i+1] = &Gem{ID: i + 1, Point: Point{X: x, Y: y}}
		}
		sc.Scan()
		opCnt, _ := strconv.Atoi(sc.Text())
		for i := 0; i < opCnt; i++ {
			sc.Scan()
			chunks := strings.Split(sc.Text(), " ")
			dir := chunks[0]
			v, _ := strconv.Atoi(chunks[1])
			m.Move(dir, v)
		}
		msg := "No"
		if m.HasCompleted() {
			msg = "Yes"
		}
		fmt.Println(msg)
	}
}
