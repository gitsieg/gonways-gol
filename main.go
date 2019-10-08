package main

import (
	"github.com/gorilla/mux"
	"gocv.io/x/gocv"
	"gonways-gol/gol"
	"gonways-gol/serve"
	"image"
	"log"
	"net/http"

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

	r := createRouter(serve.NewGameController(board))
	r.Handle("/", stream)
	log.Fatal(http.ListenAndServe(":8080", r)) // reading the stream is handled on main 'thread'.
}

func createRouter(routes... serve.Routable) *mux.Router  {
	r := mux.NewRouter()
	for _, v := range routes {
		for path, route := range v.Routes() {
			r.Handle(path, route.Handler).Methods(route.Methods...)
		}
	}
	return r
}

