package serve

import (
	"github.com/hybridgroup/mjpeg"
	"gocv.io/x/gocv"
)

func Stream(stream *mjpeg.Stream, mats <-chan gocv.Mat) {
	for {
		mat := <-mats
		buf, _ := gocv.IMEncode(".jpg", mat)
		stream.UpdateJPEG(buf)
		mat.Close()
	}
}
