package main

import (
	"image"
	"log"
	"time"
)

func main() {
	b := NewBoard(image.Pt(20, 20))
	for {
		b.Pretty()
		b.Iterate()
		time.Sleep(100 * time.Millisecond)
	}
}

// NewBoard creates a new board with given dimensions.
func NewBoard(dims image.Point) *Board {
	board := &Board{
		dims:  dims,
		cells: map[image.Point]bool{},
	}
	board.init()
	return board
}

// Board is the board with cells and dimensions
type Board struct {
	dims  image.Point
	cells map[image.Point]bool
}

// init set an inital board state.
func (b *Board) init() {
	b.cells[image.Pt(1, 0)] = true
	b.cells[image.Pt(2, 1)] = true
	b.cells[image.Pt(0, 2)] = true
	b.cells[image.Pt(1, 2)] = true
	b.cells[image.Pt(2, 2)] = true
}

// Event creates an event containing relevant board data.
func (b *Board) Event() *Event {
	// dont care about syncing right now
	return &Event{Points: b.Points()}
}

// Points returns the dimensions of the board and a list of living cells within the board
func (b *Board) Points() []image.Point {
	pts := make([]image.Point, 0)
	for k := range b.cells {
		pts = append(pts, k)
	}
	return pts
}

// Event encapsulates the board data at a given point in time.
type Event struct {
	Dims   image.Point   `json:"dims"`
	Points []image.Point `json:"points"`
}

// Pretty just prints the board to the console.
func (b *Board) Pretty() {
	out := "\n"
	for r := 0; r < b.dims.Y; r++ {
		for c := 0; c < b.dims.X; c++ {
			if b.cells[image.Pt(c, r)] {
				out += " 0 "
				continue
			}
			out += " - "
		}
		out += "\n"
	}
	log.Println(out)
}

// Iterate iterates the board.
func (b *Board) Iterate() {
	nexIteration := map[image.Point]bool{}
	for r := 0; r < b.dims.Y; r++ {
		for c := 0; c < b.dims.X; c++ {
			pt := image.Pt(c, r)
			nCount := b.neighbors(pt)
			if !b.cells[pt] && nCount == 3 {
				nexIteration[pt] = true
				continue
			}
			if b.cells[pt] && (nCount == 2 || nCount == 3) {
				nexIteration[pt] = true
				continue
			}
		}
	}
	b.cells = nexIteration
}

// neighbors counts how many neighbors a given point has.
func (b *Board) neighbors(pt image.Point) int {
	rect := image.Rect(pt.X-1, pt.Y-1, pt.X+2, pt.Y+2)
	nCount := 0
	for r := rect.Min.Y; r < rect.Max.Y; r++ {
		for c := rect.Min.X; c < rect.Max.X; c++ {
			if c == pt.X && pt.Y == r {
				continue
			}
			npt := image.Pt(c, r)
			if b.cells[npt] {
				nCount++
			}
		}
	}
	return nCount
}
