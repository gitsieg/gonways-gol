package gol

import (
	"image"
	"log"
)

// GameOfLife creates a new board with given dimensions.
func GameOfLife(dims image.Point) *Board {
	board := &Board{
		Dims:  dims,
		cells: map[image.Point]bool{},
	}
	board.init()
	return board
}
func HighLife(dims image.Point) *Board {
	board := &Board{
		highLife: true,
		Dims:     dims,
		cells:    map[image.Point]bool{},
	}
	board.init()
	return board
}

// Board is the board with cells and dimensions
type Board struct {
	highLife bool
	Dims     image.Point
	cells    map[image.Point]bool
}

// init set an inital board state.
func (b *Board) init() {

	//b.smallExploderAt(image.Pt(30, 30))
	////b.exploderAt(image.Pt(60, 60))
	//b.gosperGliderGunAt(image.Pt(0, 4))
	b.gosperGliderGunAt(image.Pt(0, 34))
	//b.gosperGliderGunAt(image.Pt(0, 64))
	//b.gosperGliderGunAt(image.Pt(0, 94))
	//b.gosperGliderGunAt(image.Pt(0, 124))
	//b.gosperGliderGunAt(image.Pt(0, 154))
	//b.gosperGliderGunAt(image.Pt(0, 184))
	//b.gosperGliderGunAt(image.Pt(0, 214))
	//b.gosperGliderGunAt(image.Pt(38, 4))
	//b.gosperGliderGunAt(image.Pt(80, 4))
	//
	//b.tumblerAt(image.Pt(300, 100))
	//b.tumblerAt(image.Pt(300, 110))
	//b.tumblerAt(image.Pt(300, 120))
	//b.tumblerAt(image.Pt(300, 130))
	//b.tumblerAt(image.Pt(300, 140))
	//b.tumblerAt(image.Pt(300, 150))
	//b.tumblerAt(image.Pt(300, 160))
	//b.tumblerAt(image.Pt(300, 170))
	//b.tumblerAt(image.Pt(300, 180))
	//b.tumblerAt(image.Pt(300, 190))
	//b.tumblerAt(image.Pt(300, 200))
	//b.tumblerAt(image.Pt(300, 210))
	//b.tumblerAt(image.Pt(300, 220))
}

func (b *Board) replicator(pt image.Point) {
	b.cells[image.Pt(pt.X+2, pt.Y)] = true
	b.cells[image.Pt(pt.X+3, pt.Y)] = true
	b.cells[image.Pt(pt.X+4, pt.Y)] = true

	b.cells[image.Pt(pt.X+1, pt.Y+1)] = true
	b.cells[image.Pt(pt.X+4, pt.Y+1)] = true

	b.cells[image.Pt(pt.X, pt.Y+2)] = true
	b.cells[image.Pt(pt.X+4, pt.Y+2)] = true

	b.cells[image.Pt(pt.X, pt.Y+3)] = true
	b.cells[image.Pt(pt.X+3, pt.Y+3)] = true

	b.cells[image.Pt(pt.X, pt.Y+4)] = true
	b.cells[image.Pt(pt.X+1, pt.Y+4)] = true
	b.cells[image.Pt(pt.X+2, pt.Y+4)] = true
}

/*
  0 1 2 3 4
0 - - o o o
1 - o - - o
2 o - - - o
3 o - - o
4 o o o

*/

func (b *Board) smallExploderAt(pt image.Point) {
	b.cells[image.Pt(pt.X+1, pt.Y+0)] = true
	b.cells[image.Pt(pt.X+0, pt.Y+1)] = true
	b.cells[image.Pt(pt.X+1, pt.Y+1)] = true
	b.cells[image.Pt(pt.X+2, pt.Y+1)] = true
	b.cells[image.Pt(pt.X+0, pt.Y+2)] = true
	b.cells[image.Pt(pt.X+2, pt.Y+2)] = true
	b.cells[image.Pt(pt.X+1, pt.Y+3)] = true
}

func (b *Board) exploderAt(pt image.Point) {
	b.cells[image.Pt(pt.X+0, pt.Y+0)] = true
	b.cells[image.Pt(pt.X+0, pt.Y+1)] = true
	b.cells[image.Pt(pt.X+0, pt.Y+2)] = true
	b.cells[image.Pt(pt.X+0, pt.Y+3)] = true
	b.cells[image.Pt(pt.X+0, pt.Y+4)] = true
	b.cells[image.Pt(pt.X+2, pt.Y+0)] = true
	b.cells[image.Pt(pt.X+2, pt.Y+4)] = true
	b.cells[image.Pt(pt.X+4, pt.Y+0)] = true
	b.cells[image.Pt(pt.X+4, pt.Y+1)] = true
	b.cells[image.Pt(pt.X+4, pt.Y+2)] = true
	b.cells[image.Pt(pt.X+4, pt.Y+3)] = true
	b.cells[image.Pt(pt.X+4, pt.Y+4)] = true
}

func (b *Board) gliderAt(pt image.Point) {
	b.cells[image.Pt(pt.X+1, pt.Y+0)] = true
	b.cells[image.Pt(pt.X+2, pt.Y+1)] = true
	b.cells[image.Pt(pt.X+0, pt.Y+2)] = true
	b.cells[image.Pt(pt.X+1, pt.Y+2)] = true
	b.cells[image.Pt(pt.X+2, pt.Y+2)] = true
}

func (b *Board) tenCellRow(pt image.Point) {
	b.cells[image.Pt(pt.X, pt.Y)] = true
	b.cells[image.Pt(pt.X+1, pt.Y)] = true
	b.cells[image.Pt(pt.X+2, pt.Y)] = true
	b.cells[image.Pt(pt.X+3, pt.Y)] = true
	b.cells[image.Pt(pt.X+4, pt.Y)] = true
	b.cells[image.Pt(pt.X+5, pt.Y)] = true
	b.cells[image.Pt(pt.X+6, pt.Y)] = true
	b.cells[image.Pt(pt.X+7, pt.Y)] = true
	b.cells[image.Pt(pt.X+8, pt.Y)] = true
	b.cells[image.Pt(pt.X+9, pt.Y)] = true
}

func (b *Board) tumblerAt(pt image.Point) {
	b.cells[image.Pt(pt.X+1, pt.Y)] = true
	b.cells[image.Pt(pt.X+2, pt.Y)] = true
	b.cells[image.Pt(pt.X+4, pt.Y)] = true
	b.cells[image.Pt(pt.X+5, pt.Y)] = true

	b.cells[image.Pt(pt.X+1, pt.Y+1)] = true
	b.cells[image.Pt(pt.X+2, pt.Y+1)] = true
	b.cells[image.Pt(pt.X+4, pt.Y+1)] = true
	b.cells[image.Pt(pt.X+5, pt.Y+1)] = true

	b.cells[image.Pt(pt.X+2, pt.Y+2)] = true
	b.cells[image.Pt(pt.X+4, pt.Y+2)] = true

	b.cells[image.Pt(pt.X, pt.Y+3)] = true
	b.cells[image.Pt(pt.X+2, pt.Y+3)] = true
	b.cells[image.Pt(pt.X+4, pt.Y+3)] = true
	b.cells[image.Pt(pt.X+6, pt.Y+3)] = true

	b.cells[image.Pt(pt.X, pt.Y+4)] = true
	b.cells[image.Pt(pt.X+2, pt.Y+4)] = true
	b.cells[image.Pt(pt.X+4, pt.Y+4)] = true
	b.cells[image.Pt(pt.X+6, pt.Y+4)] = true

	b.cells[image.Pt(pt.X, pt.Y+5)] = true
	b.cells[image.Pt(pt.X+1, pt.Y+5)] = true
	b.cells[image.Pt(pt.X+5, pt.Y+5)] = true
	b.cells[image.Pt(pt.X+6, pt.Y+5)] = true
}

func (b *Board) gosperGliderGunAt(pt image.Point) {
	b.cells[image.Pt(pt.X+23, pt.Y)] = true
	b.cells[image.Pt(pt.X+24, pt.Y)] = true
	b.cells[image.Pt(pt.X+34, pt.Y)] = true
	b.cells[image.Pt(pt.X+35, pt.Y)] = true

	b.cells[image.Pt(pt.X+22, pt.Y+1)] = true
	b.cells[image.Pt(pt.X+24, pt.Y+1)] = true
	b.cells[image.Pt(pt.X+34, pt.Y+1)] = true
	b.cells[image.Pt(pt.X+35, pt.Y+1)] = true

	b.cells[image.Pt(pt.X, pt.Y+2)] = true
	b.cells[image.Pt(pt.X+1, pt.Y+2)] = true
	b.cells[image.Pt(pt.X+9, pt.Y+2)] = true
	b.cells[image.Pt(pt.X+10, pt.Y+2)] = true
	b.cells[image.Pt(pt.X+22, pt.Y+2)] = true
	b.cells[image.Pt(pt.X+23, pt.Y+2)] = true

	b.cells[image.Pt(pt.X, pt.Y+3)] = true
	b.cells[image.Pt(pt.X+1, pt.Y+3)] = true
	b.cells[image.Pt(pt.X+8, pt.Y+3)] = true
	b.cells[image.Pt(pt.X+10, pt.Y+3)] = true

	b.cells[image.Pt(pt.X+8, pt.Y+4)] = true
	b.cells[image.Pt(pt.X+9, pt.Y+4)] = true
	b.cells[image.Pt(pt.X+16, pt.Y+4)] = true
	b.cells[image.Pt(pt.X+17, pt.Y+4)] = true

	b.cells[image.Pt(pt.X+16, pt.Y+5)] = true
	b.cells[image.Pt(pt.X+18, pt.Y+5)] = true

	b.cells[image.Pt(pt.X+16, pt.Y+6)] = true

	b.cells[image.Pt(pt.X+35, pt.Y+7)] = true
	b.cells[image.Pt(pt.X+36, pt.Y+7)] = true

	b.cells[image.Pt(pt.X+35, pt.Y+8)] = true
	b.cells[image.Pt(pt.X+37, pt.Y+8)] = true

	b.cells[image.Pt(pt.X+35, pt.Y+9)] = true

	b.cells[image.Pt(pt.X+24, pt.Y+12)] = true
	b.cells[image.Pt(pt.X+25, pt.Y+12)] = true
	b.cells[image.Pt(pt.X+26, pt.Y+12)] = true

	b.cells[image.Pt(pt.X+24, pt.Y+13)] = true

	b.cells[image.Pt(pt.X+25, pt.Y+14)] = true
}

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
		out <- b.Points() // this will throttle, most definitely.
		b.Iterate()
		//time.Sleep(3 * time.Second)
		// if fps should be adjusted, this method is probably the place to do it
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
	if b.highLife {
		if !alive && neighbors == 3 || !alive && neighbors == 6 {
			return true
		}
	}
	if !alive && neighbors == 3 {
		return true
	}
	return alive && (neighbors == 2 || neighbors == 3)
}
