package serve

import (
	"image"

	"gocv.io/x/gocv"
)

func MatProduce(dims image.Point, in <-chan []image.Point, out chan<- gocv.Mat) {
	orig := gocv.NewMatWithSize(dims.Y, dims.X, gocv.MatTypeCV8UC1)
	defer orig.Close()
	for {
		points := <-in
		// violates resource release convention. as of now handled in stream
		src := gocv.NewMat()
		orig.CopyTo(&src)
		for _, pt := range points {
			src.SetUCharAt(pt.Y, pt.X, 255)
		}
		gocv.Resize(src, &src, image.Pt(640, 480), 0, 0, gocv.InterpolationNearestNeighbor)
		out <- src
	}
}
