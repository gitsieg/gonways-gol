package gol

import (
	"errors"
	"fmt"
	"image"
	"log"
	"sync"
)

// GameOfLife creates a new board with given dimensions.
func GameOfLife(dims image.Point) *Board {
	board := &Board{
		Dims:     dims,
		cells:    map[image.Point]bool{},
		requests: make(chan *StructuralRequest, 0),
	}
	board.init()
	return board
}

// Board is the board with cells and dimensions
type Board struct {
	mu       sync.Mutex
	highLife bool
	Dims     image.Point
	cells    map[image.Point]bool
	requests chan *StructuralRequest
}

//
// Event creates an event containing relevant board data.
func (b *Board) Event() *Event {
	// dont care about syncing right now
	return &Event{
		Dims:   b.Dims,
		Points: b.Points(),
	}
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
	Dims   image.Point   `json:"Dims"`
	Points []image.Point `json:"points"`
}

// Pretty just prints the board to the console.
func (b *Board) Pretty() {
	out := "\n"
	for r := 0; r < b.Dims.Y; r++ {
		for c := 0; c < b.Dims.X; c++ {
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
	for r := 0; r < b.Dims.Y; r++ {
		for c := 0; c < b.Dims.X; c++ {
			pt := image.Pt(c, r)
			if b.continueLivingOrResurrect(b.cells[pt], b.neighbors(pt)) {
				nexIteration[pt] = true
			}
		}
	}
	b.cells = nexIteration
}

// Runs the board in a separate routine forever.
func (b *Board) Start(out chan<- []image.Point) {
	for {
		b.mu.Lock()
		out <- b.Points()
		b.Iterate()
		b.mu.Unlock()
	}
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

// continueLivingOrResurrect returns true if the cell should be alive, false otherwise.
func (b *Board) continueLivingOrResurrect(alive bool, neighbors int) bool {
	if !alive && neighbors == 3 {
		return true
	}
	return alive && (neighbors == 2 || neighbors == 3)
}

// Clear removes all alive cells from the board.
func (b *Board) Clear() {
	b.mu.Lock()
	b.cells = map[image.Point]bool{}
	defer b.mu.Unlock()
}

// Handle requests board
func (b *Board) Handle(r StructuralRequest) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	if !r.At().In(image.Rect(0, 0, b.Dims.X, b.Dims.Y)) {
		return fmt.Errorf("request point %v out of bounds", r)
	}
	patternFunc, e := MapToFunc(r.Type())
	if e != nil {
		return e
	}
	patternFunc(b, r.At())
	return nil
}

func (b *Board) init() {

}

type StructuralRequest interface {
	At() image.Point
	Type() GolPattern
}

func MapToFunc(r GolPattern) (StructFunc, error) {
	switch r {
	case Tumbler:
		return TumblerAt, nil
	case GosperGlider:
		return GliderGunAt, nil
	case Replicator:
		return ReplicatorAt, nil
	case SmallExploder:
		return SmallExploderAt, nil
	case Exploder:
		return ExploderAt, nil
	case Glider:
		return GliderAt, nil
	case TenCellRow:
		return TenCellRowAt, nil
	case Clear:
		return ClearBoard, nil
	default:
		return nil, errors.New("no such pattern type")
	}
}

type StructFunc func(b *Board, pt image.Point)

type GolPattern int

func (g GolPattern) String() string {
	return [...]string{
		"Tumbler",
		"GosperGlider",
		"Replicator",
		"SmallExploder",
		"Exploder",
		"Glider",
		"TenCellRow",
		"Clear",
	}[g]
}

const (
	Tumbler GolPattern = iota
	GosperGlider
	Replicator
	SmallExploder
	Exploder
	Glider
	TenCellRow
	Clear
)
