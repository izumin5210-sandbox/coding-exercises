package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

type Color int

const (
	ColorBlack Color = iota
	ColorRed
)

var strByColor = map[Color]string{
	ColorBlack: "black",
	ColorRed:   "red",
}

func (c Color) String() string {
	return strByColor[c]
}

type Node struct {
	ID      int
	X, Y    int
	Color   Color
	NextIDs map[int]struct{}
}

func (n *Node) String() string {
	d, _ := json.Marshal(n)
	return string(d)
}

type Map struct {
	Width, Height int
	Nodes         map[int]*Node
}

func (m *Map) GetByID(id int) *Node {
	return m.Nodes[id]
}

func (m *Map) Get(x, y int) *Node {
	return m.Nodes[m.getID(x, y)]
}

func (m *Map) SetColor(x, y int, c Color) {
	m.Nodes[m.getID(x, y)].Color = c
}

func (m *Map) getID(x, y int) int {
	return x + y*m.Width
}

func (m *Map) init() {
	for x := 1; x <= m.Width; x++ {
		for y := 1; y <= m.Height; y++ {
			id := m.getID(x, y)
			m.Nodes[id] = &Node{ID: id, X: x, Y: y}
		}
	}
	for x := 1; x <= m.Width; x++ {
		for y := 1; y <= m.Height; y++ {
			nextIDs := map[int]struct{}{}
			if x > 1 {
				nextIDs[m.getID(x-1, y)] = struct{}{}
			}
			if x < m.Width {
				nextIDs[m.getID(x+1, y)] = struct{}{}
			}
			if y > 1 {
				nextIDs[m.getID(x, y-1)] = struct{}{}
			}
			if y < m.Height {
				nextIDs[m.getID(x, y+1)] = struct{}{}
			}
			m.Get(x, y).NextIDs = nextIDs
		}
	}
}

func NewMap(w, h int) *Map {
	m := &Map{Width: w, Height: h, Nodes: map[int]*Node{}}
	m.init()
	return m
}

type Searcher struct {
	Map        *Map
	queue      []*Node
	touchedIDs map[int]struct{}
	visitedIDs map[int]struct{}
}

func NewSearcher(m *Map) *Searcher {
	return &Searcher{
		Map:        m,
		queue:      []*Node{},
		touchedIDs: map[int]struct{}{},
		visitedIDs: map[int]struct{}{},
	}
}

func (s *Searcher) Search(start *Node) {
	s.enqueue(start)

	for {
		if len(s.queue) == 0 {
			break
		}
		for nextID := range s.dequeue().NextIDs {
			n := s.Map.GetByID(nextID)
			if n.Color == ColorRed || s.hasTouched(n) {
				continue
			}
			s.enqueue(n)
		}
	}
}

func (s *Searcher) VisitedNodeCount() int {
	return len(s.visitedIDs)
}

func (s *Searcher) hasTouched(n *Node) bool {
	_, ok := s.touchedIDs[n.ID]
	return ok
}

func (s *Searcher) enqueue(n *Node) {
	s.queue = append(s.queue, n)
	s.touchedIDs[n.ID] = struct{}{}
}

func (s *Searcher) dequeue() *Node {
	n := s.queue[0]
	s.queue = s.queue[1:]
	s.visitedIDs[n.ID] = struct{}{}
	return n
}

func main() {
	for sc.Scan() {
		if sc.Text() == "0 0" {
			break
		}
		chunks := strings.Split(sc.Text(), " ")
		w, _ := strconv.Atoi(chunks[0])
		h, _ := strconv.Atoi(chunks[1])

		m := NewMap(w, h)

		// Set colors
		var start *Node
		for y := 1; y <= h; y++ {
			sc.Scan()
			for i, c := range sc.Text() {
				switch c {
				case '.', '@':
					m.SetColor(i+1, y, ColorBlack)
					if c == '@' {
						start = m.Get(i+1, y)
					}
				case '#':
					m.SetColor(i+1, y, ColorRed)
				}
			}
		}

		// Search
		s := NewSearcher(m)
		s.Search(start)
		fmt.Println(s.VisitedNodeCount())
	}
}
