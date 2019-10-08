package serve

import (
	"log"
	"net/http"
)

// Middleware
type cqrs struct {
	headerOptions map[string]string
}

func (c *cqrs) Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			for k, v := range c.headerOptions {
				w.Header().Set(k, v)
			}
			log.Println("Options")
			return
		}
		for k, v := range c.headerOptions {
			w.Header().Set(k, v)
		}
		log.Println("Options")
		h.ServeHTTP(w, r)
		log.Println("middleware end")
	})
}

func ContentJSON() func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("content-type", "application/json")
			next.ServeHTTP(w, r)
		})
	}
}

func Headers() func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println(len(w.Header()))
			log.Println(w.Header())
			next.ServeHTTP(w, r)
		})
	}
}
