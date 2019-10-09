package serve

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Response is the response written to the http response.
type Response interface {
	WriteTo(w http.ResponseWriter) (n int64, err error)
	Status() int
}

type resp struct {
	message    string
	statusCode int
	err        error
}

// WriteTo writes the resp status code onto the headers of the http response, and any message or error associated with it.
func (r resp) WriteTo(w http.ResponseWriter) (n int64, err error) {
	if r.err != nil {
		i, err := fmt.Fprintf(w, r.err.Error())
		return int64(i), err
	}
	i, err := fmt.Fprintf(w, r.message)
	return int64(i), err
}

// Status returns the resp StatusCode
func (r resp) Status() int {
	return r.statusCode
}

// StatusOK returns a StatusOK response
func StatusOK() Response {
	return &resp{statusCode: http.StatusOK}
}

// StatusOK returns a StatusOK response
func StatusNoContent() Response {
	return &resp{statusCode: http.StatusNoContent}
}

func InternalServerError() Response {
	return &resp{statusCode: http.StatusInternalServerError}
}

func BadRequest(err error) Response {
	return &resp{
		statusCode: http.StatusBadRequest,
		err:        err,
	}
}

// ControllerFunc is the interface to be implemented by Controllers
type ControllerFunc func(ctx Context) Response

// Handler creates a http.HandlerFunc from the given ControllerFunc
func Handler(fn ControllerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handler start")
		ctx := Context{
			Request:  r,
			Response: w,
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//"Access-Control-Allow-Origin" : "*",
		fn(ctx).WriteTo(w)
		log.Println("Handler end")
	}
}

// Context wraps the http request and response
type Context struct {
	Request  *http.Request
	Response http.ResponseWriter
}

func (c Context) Deadline() (deadline time.Time, ok bool) {
	return c.Request.Context().Deadline()
}

func (c Context) Done() <-chan struct{} {
	return c.Request.Context().Done()
}

func (c Context) Err() error {
	return c.Request.Context().Err()
}

func (c Context) Value(key interface{}) interface{} {
	return c.Request.Context().Value(key)
}

// Routable is implemented by controllers
type Routable interface {
	// Routes is the available routes for implementors.
	Routes() Routes
}

// Path is the path
type Path = string

// Routes is the mapping of a path to a Route
type Routes map[Path]Route

// Methods is the http methods accepted by a route
type Methods = []string

// Route is the endpoint for a request with methods.
type Route struct {
	Methods Methods
	Handler http.Handler
}
