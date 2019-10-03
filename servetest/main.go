package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/time", TimerHandler()(timeHandler{format: time.RFC1123}))
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func TimerHandler() func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t1 := time.Now()
			next.ServeHTTP(w, r)
			t2 := time.Now()
			log.Printf("%v", t2.Sub(t1))
		})
	}
}

type timeHandler struct {
	format string
}

func (h timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(h.format)
	fmt.Fprintf(w, tm)
}
