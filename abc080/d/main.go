package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Program struct {
	Start   float64 `json:"start"`
	End     float64 `json:"end"`
	Channel int     `json:"channel"`
}

type Programs []*Program

func (p Programs) Len() int {
	return len(p)
}

func (p Programs) Less(i, j int) bool {
	return p[i].Start < p[j].Start
}

func (p Programs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *Program) String() string {
	d, _ := json.Marshal(p)
	return string(d)
}

type Recorder struct {
	ID      int
	Current *Program
}

func (r *Recorder) CanRecord(p *Program) bool {
	start := p.Start
	if r.Current.Channel != p.Channel {
		start -= 0.5
		return r.Current.End < start
	}
	return r.Current.End <= start
}

func (r *Recorder) String() string {
	return fmt.Sprintf("recorder:%d", r.ID)
}

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Scan()
	row := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(row[0])
	// c, _ := strconv.Atoi(row[1])
	programs := make([]*Program, n, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		row = strings.Split(sc.Text(), " ")
		s, _ := strconv.ParseFloat(row[0], 64)
		t, _ := strconv.ParseFloat(row[1], 64)
		c, _ := strconv.Atoi(row[2])
		programs[i] = &Program{
			Start:   s,
			End:     t,
			Channel: c,
		}
	}

	sort.Sort(Programs(programs))

	recorders := []*Recorder{}

	for _, p := range programs {
		var nextRecorders []*Recorder
		for _, r := range recorders {
			if r.CanRecord(p) {
				nextRecorders = append(nextRecorders, r)
			}
		}
		if len(nextRecorders) == 0 {
			r := &Recorder{ID: len(recorders) + 1}
			recorders = append(recorders, r)
			nextRecorders = append(nextRecorders, r)
		}
		var nextRecorder *Recorder
		// if len(nextRecorders) > 1 {
		//   for i, r := range nextRecorders {
		//     if r.Current.Channel == p.Channel {
		//       nextRecorder = r
		//       break
		//     }
		//     skipped := false
		//     for j := 1; j <= c && j < n; j++ {
		//       if programs[i+j].Channel == r.Current.Channel && p.End >= programs[j+j].Start-0.5 {
		//         skipped = true
		//         break
		//       }
		//     }
		//     if !skipped {
		//       nextRecorder = r
		//       break
		//     }
		//   }
		// }
		if nextRecorder == nil {
			nextRecorder = nextRecorders[0]
		}
		nextRecorder.Current = p
	}

	fmt.Println(len(recorders))
}
