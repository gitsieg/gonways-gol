package gol

import (
	"image"
)

// todo: restructure to receiver b.
func ReplicatorAt(b *Board, pt image.Point) {
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

func SmallExploderAt(b *Board, pt image.Point) {
	b.cells[image.Pt(pt.X+1, pt.Y+0)] = true
	b.cells[image.Pt(pt.X+0, pt.Y+1)] = true
	b.cells[image.Pt(pt.X+1, pt.Y+1)] = true
	b.cells[image.Pt(pt.X+2, pt.Y+1)] = true
	b.cells[image.Pt(pt.X+0, pt.Y+2)] = true
	b.cells[image.Pt(pt.X+2, pt.Y+2)] = true
	b.cells[image.Pt(pt.X+1, pt.Y+3)] = true
}

func ExploderAt(b *Board, pt image.Point) {
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

func GliderAt(b *Board, pt image.Point) {
	b.cells[image.Pt(pt.X+1, pt.Y+0)] = true
	b.cells[image.Pt(pt.X+2, pt.Y+1)] = true
	b.cells[image.Pt(pt.X+0, pt.Y+2)] = true
	b.cells[image.Pt(pt.X+1, pt.Y+2)] = true
	b.cells[image.Pt(pt.X+2, pt.Y+2)] = true
}

func TenCellRowAt(b *Board, pt image.Point) {
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

func ClearBoard(b *Board, pt image.Point)  {
	b.cells = map[image.Point]bool{}
}

func TumblerAt(b *Board, pt image.Point) {
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

func GliderGunAt(b *Board, pt image.Point) {
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
