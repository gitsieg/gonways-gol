package server

import (
	"encoding/json"
	"github.com/rs/cors"
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
	c := cors.New(cors.Options{
		AllowedOrigins:         []string{"*"},
		AllowOriginFunc:        nil,
		AllowOriginRequestFunc: nil,
		AllowedMethods:         nil,
		AllowedHeaders:         nil,
		ExposedHeaders:         nil,
		MaxAge:                 0,
		AllowCredentials:       false,
		OptionsPassthrough:     false,
		Debug:                  false,
	})
	c.Handler(server.httpServe.Handler)
	return server
}

func (s *Server) HandleBoardCreateStructure() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &gol.GolRequest{}
		decoder := json.NewDecoder(r.Body)
		if e := decoder.Decode(req); e != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(e.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
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

