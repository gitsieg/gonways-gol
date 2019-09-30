package main

import (
	"gonways-gol/gol"
	"gonways-gol/server"
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
	serv := server.NewServer(board)

	// Set ut parallelism
	go serv.Start()
	go board.Start(points)                   // board feeds to points
	go server.MatProduce(size, points, mats) // matproduce reads from points and produces mats
	go server.Stream(stream, mats)           // stream reads from mats and produces to jpeg buffer
	http.Handle("/", stream)
	log.Fatal(http.ListenAndServe(":8080", nil)) // reading the stream is handled on main 'thread'.
}
