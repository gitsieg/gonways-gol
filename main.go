package main

import (
	"gonways-gol/gol"
	"gonways-gol/serve"
	"image"
	"log"
	"net/http"

	"gocv.io/x/gocv"

	"github.com/hybridgroup/mjpeg"
)

func main() {
	size := image.Pt(320, 240)
	points := make(chan []image.Point)
	mats := make(chan gocv.Mat)
	stream := mjpeg.NewStream()
	board := gol.GameOfLife(size)

	// Set ut parallelism
	go board.Start(points)                  // board feeds to points
	go serve.MatProduce(size, points, mats) // matproduce reads from points and produces mats
	go serve.Stream(stream, mats)           // stream reads from mats and produces to jpeg buffer

	//http.Handle("/", stream)
	http.Handle("/", SetupHandlers(serve.NewGameController(board)))
	log.Fatal(http.ListenAndServe(":8080", nil)) // reading the stream is handled on main 'thread'.
}

func SetupHandlers(controller *serve.GameController) *http.ServeMux {
	mux := http.NewServeMux()
	for k, v := range controller.Routes() {
		mux.HandleFunc(k, v)
	}
	//mux.Handle("/stream", &stream)
	return mux
}
