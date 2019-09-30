package gol

import (
	"image"
	"reflect"
	"testing"
)

func TestBoard_neighbors(t *testing.T) {
	tests := []struct {
		name    string
		Dims    image.Point
		cells   map[image.Point]bool
		checkPt image.Point
		want    int
	}{
		{
			name: "test 3x3 check center 5 neighbors",
			Dims: image.Pt(3, 3),
			cells: map[image.Point]bool{
				image.Pt(0, 0): true,
				image.Pt(1, 0): true,
				image.Pt(2, 0): true,
				image.Pt(0, 1): true,
				image.Pt(2, 1): true,
			},
			checkPt: image.Pt(1, 1),
			want:    5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				highLife: false,
				Dims:     tt.Dims,
				cells:    tt.cells,
			}
			if got := b.neighbors(tt.checkPt); got != tt.want {
				t.Errorf("Board.neighbors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_Iterate(t *testing.T) {
	tests := []struct {
		name    string
		dims    image.Point
		initial map[image.Point]bool
		want    map[image.Point]bool
	}{
		{
			dims: image.Pt(3, 3),
			initial: map[image.Point]bool{
				image.Pt(0, 0): true,
				image.Pt(1, 0): true,
				image.Pt(2, 0): true,
			},
			want: map[image.Point]bool{
				image.Pt(1, 0): true,
				image.Pt(1, 1): true,
			},
		},
		{
			dims: image.Pt(3, 3),
			initial: map[image.Point]bool{
				image.Pt(1, 0): true,
				image.Pt(1, 1): true,
				image.Pt(1, 2): true,
			},
			want: map[image.Point]bool{
				image.Pt(0, 1): true,
				image.Pt(1, 1): true,
				image.Pt(2, 1): true,
			},
		},
		{
			dims: image.Pt(3, 3),
			initial: map[image.Point]bool{
				image.Pt(0, 0): true,
				image.Pt(2, 0): true,
				image.Pt(0, 2): true,
				image.Pt(2, 2): true,
			},
			want: map[image.Point]bool{},
		},
		{
			dims: image.Pt(3, 3),
			initial: map[image.Point]bool{
				image.Pt(0, 0): true,
				image.Pt(2, 0): true,
				image.Pt(1, 1): true,
				image.Pt(0, 2): true,
				image.Pt(2, 2): true,
			},
			want: map[image.Point]bool{
				image.Pt(1, 0): true,
				image.Pt(0, 1): true,
				image.Pt(2, 1): true,
				image.Pt(1, 2): true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				Dims:  tt.dims,
				cells: tt.initial,
			}
			b.Iterate()
			if !reflect.DeepEqual(b.cells, tt.want) {
				t.Fail()
			}
		})
	}
}
