package serve

import (
	"fmt"
	"gonways-gol/gol"
	"image"
	"net/http"
	"time"
)

type ControllerFunc func(ctx Context) Response

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

type Response interface {
	WriteTo(w http.ResponseWriter) (n int64, err error)
	Status() int
}

func Handler(fn ControllerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := Context{
			Request:  r,
			Response: w,
		}
		fn(ctx).WriteTo(w)
	}
}

func NewGameController(game *gol.Board) *GameController {
	c := &GameController{
		game: game,
	}
	c.init()
	return c
}
func (c *GameController) init() {
	c.routes = c.Routes()
}

func (c *GameController) Routes() Routes {
	return Routes{
		"game/clear":  Handler(c.HandleBoardClear),
		"game/create": Handler(c.HandleBoardCreate),
	}
}

type GameController struct {
	game   *gol.Board
	routes Routes
}

type Routes map[string]http.HandlerFunc

func (c *GameController) HandleBoardClear(ctx Context) Response {
	c.game.Request(&gol.BoardRequest{
		Request: gol.ClearBoard,
	})
	return OK()
}

func (c *GameController) HandleBoardCreate(ctx Context) Response {
	c.game.Request(&gol.BoardRequest{
		Request: gol.Tumbler,
		At:      image.Pt(100, 100),
	})
	return OK()
}

func OK() Response {
	return &StatusOK{}
}

type StatusOK struct {
}

func (o StatusOK) WriteTo(w http.ResponseWriter) (n int64, err error) {
	w.WriteHeader(o.Status())
	i, err := fmt.Fprintf(w, "OK")
	return int64(i), err
}

func (o StatusOK) Status() int {
	return http.StatusOK
}
