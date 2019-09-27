package main

import (
	"gonways-gol/gol"
	"image"
	"log"
	"net/http"
	"time"

	"gocv.io/x/gocv"

	"github.com/hybridgroup/mjpeg"
)

var board *gol.Board
var stream *mjpeg.Stream

func main() {
	stream = mjpeg.NewStream()
	pt := image.Pt(320, 240)
	board = gol.NewBoard(pt)
	go mjpegTransformAndStream()
	http.Handle("/", stream)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mjpegTransformAndStream() {
	orig := gocv.NewMatWithSize(240, 320, gocv.MatTypeCV8UC1)
	defer orig.Close()
	for {
		img := gocv.NewMat()
		orig.CopyTo(&img)
		for _, v := range board.Points() {
			img.SetUCharAt(v.Y, v.X, 255)
		}
		gocv.Resize(img, &img, image.Pt(640, 480), 0, 0, gocv.InterpolationNearestNeighbor)
		//gocv.Resize(img, &img, image.Pt(1920, 1080), 0, 0, gocv.InterpolationNearestNeighbor)
		buf, _ := gocv.IMEncode(".jpg", img)
		stream.UpdateJPEG(buf)
		board.Iterate()
		img.Close()
		time.Sleep(41 * time.Millisecond)
	}
}
