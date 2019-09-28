package server

import (
	"encoding/json"
	"gonways-gol/gol"
	"image"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	httpServe *http.Server
	board     *gol.Board
}

func NewServer() *Server {
	server := &Server{
		httpServe: &http.Server{
			Addr:              "127.0.0.1:8080",
			Handler:           nil,
			TLSConfig:         nil,
			ReadTimeout:       15 * time.Second,
			ReadHeaderTimeout: 0,
			WriteTimeout:      15 * time.Second,
			IdleTimeout:       0,
			MaxHeaderBytes:    0,
			TLSNextProto:      nil,
			ConnState:         nil,
			ErrorLog:          nil,
		},
		board: gol.GameOfLife(image.Pt(60, 60)),
	}
	router := mux.NewRouter()
	router.HandleFunc("/hello", server.handleIterate).Methods("GET")
	server.httpServe.Handler = router
	return server
}

func (s *Server) handleIterate(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	event := s.board.Event()
	bytes, e := json.Marshal(event)
	if e != nil {
		w.Write([]byte(e.Error()))
	}
	w.Write(bytes)
	s.board.Iterate()
	s.board.Pretty()
}

func (s *Server) Start() error {
	return s.httpServe.ListenAndServe()
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-Type", "application/json")
}
