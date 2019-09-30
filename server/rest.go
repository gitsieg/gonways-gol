package server

import (
	"encoding/json"
	"gonways-gol/gol"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	httpServe *http.Server
	board     *gol.Board
}

func NewServer(board *gol.Board) *Server {
	server := &Server{
		httpServe: &http.Server{
			Addr:              "127.0.0.1:8081",
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
		board: board,
	}
	server.httpServe.Handler = server.routes()
	return server
}

func (s *Server) HandleBoardCreateStructure() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		enableCors(&w)
		req := &gol.GolRequest{}
		decoder := json.NewDecoder(r.Body)
		if e := decoder.Decode(req); e != nil {

		}
		log.Println(req)
		s.board.Requests <- req
	}
}

func (s *Server) Start() error {
	return s.httpServe.ListenAndServe()
}

func (s *Server) routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/create", s.HandleBoardCreateStructure()).Methods("POST")
	return router
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-Type", "application/json")
}
