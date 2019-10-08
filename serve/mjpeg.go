package serve

import (
	"github.com/hybridgroup/mjpeg"
	"gocv.io/x/gocv"
	"image"
)

func Stream(stream *mjpeg.Stream, mats <-chan gocv.Mat) {
	for {
		mat := <-mats
		buf, _ := gocv.IMEncode(".jpg", mat)
		stream.UpdateJPEG(buf)
		mat.Close()
	}
}

// MatProduce produces grayscale mats with the given dims to the 'out' channel. The points received from 'in' will be
// be white in out mats.
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

