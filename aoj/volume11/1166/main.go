package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

type Room struct {
	ID    int
	Nexts []*Room
}

type Maze struct {
	Width, Height int
	rooms         map[int]*Room
}

func NewMaze(width, height int) *Maze {
	m := &Maze{width, height, map[int]*Room{}}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			id := m.calcID(x, y)
			m.rooms[id] = &Room{ID: id}
		}
	}
	return m
}

func (m *Maze) Connect(x1, y1, x2, y2 int) {
	r1 := m.GetRoom(x1, y1)
	r2 := m.GetRoom(x2, y2)
	r1.Nexts = append(r1.Nexts, r2)
	r2.Nexts = append(r2.Nexts, r1)
}

func (m *Maze) GetRoom(x, y int) *Room {
	return m.GetRoomByID(m.calcID(x, y))
}

func (m *Maze) GetRoomByID(id int) *Room {
	return m.rooms[id]
}

func (m *Maze) GetGoal() *Room {
	return m.GetRoom(m.Width-1, m.Height-1)
}

func (m *Maze) calcID(x, y int) int {
	return x + y*m.Width
}

type Searcher struct {
	maze      *Maze
	queue     []*Room
	distances map[int]int
	prevs     map[int]*Room
}

func NewSearcher(m *Maze) *Searcher {
	return &Searcher{maze: m}
}

func (s *Searcher) SearchShotestPath() []*Room {
	s.init()
	var current *Room
	for len(s.queue) > 0 {
		current = s.dequeue()
		d := s.distances[current.ID]
		for _, r := range current.Nexts {
			if dn, ok := s.distances[r.ID]; !ok || dn > d+1 {
				s.distances[r.ID] = d + 1
				s.prevs[r.ID] = current
			}
		}
	}
	current = s.maze.GetGoal()
	path := []*Room{current}
	for {
		current = s.prevs[current.ID]
		if current == nil {
			break
		}
		path = append([]*Room{current}, path...)
	}
	return path
}

func (s *Searcher) init() {
	s.distances = map[int]int{}
	s.prevs = map[int]*Room{}
	for x := 0; x < s.maze.Width; x++ {
		for y := 0; y < s.maze.Height; y++ {
			r := s.maze.GetRoom(x, y)
			s.queue = append(s.queue, r)
		}
	}
	start := s.maze.GetRoom(0, 0)
	s.distances[start.ID] = 0
}

func (s *Searcher) dequeue() *Room {
	var nearestIdx int
	minDistance := math.Inf(0)
	for i, r := range s.queue {
		if d, ok := s.distances[r.ID]; ok && float64(d) < minDistance {
			nearestIdx = i
		}
	}
	r := s.queue[nearestIdx]
	s.queue = append(s.queue[0:nearestIdx], s.queue[nearestIdx+1:len(s.queue)]...)
	return r
}

func main() {
	for sc.Scan() {
		if sc.Text() == "0 0" {
			break
		}
		chunks := strings.Split(sc.Text(), " ")
		w, _ := strconv.Atoi(chunks[0])
		h, _ := strconv.Atoi(chunks[1])

		m := NewMaze(w, h)

		for i := 0; i < 2*h-1; i++ {
			sc.Scan()
			for j, c := range sc.Text() {
				switch c {
				case ' ':
					// do nothing
				case '0':
					if i%2 == 0 {
						m.Connect(j/2+1, i/2, j/2, i/2)
					} else {
						m.Connect(j/2, i/2+1, j/2, i/2)
					}
				case '1':
					// do nothing
				}
			}
		}

		s := NewSearcher(m)
		path := s.SearchShotestPath()
		fmt.Println(len(path))
	}
}
