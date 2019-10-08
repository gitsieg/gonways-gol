package serve

import (
	"encoding/json"
	"gonways-gol/gol"
	"image"
	"log"
)

func NewGameController(game *gol.Board) *GameController {
	c := &GameController{
		game: game,
		cqrs:cqrs{headerOptions: map[string]string{
			"Access-Control-Allow-Origin" : "*",
			"Access-Control-Allow-Headers" : "Content-Type",
			//"Content-Type" : "application/json",
		}},
	}
	return c
}

func (c *GameController) Routes() Routes {
	withCQRS := c.cqrs.Middleware
	withJson := ContentJSON()
	return Routes{
		"/game/clear":
			Route{Methods: Methods{"POST", "OPTIONS"}, Handler:withCQRS( Handler(c.HandleBoardClear))},
		"/game/create":
			Route{Methods: Methods{"POST", "OPTIONS"}, Handler: withCQRS(Handler(c.HandleBoardCreate))},
		"/game/dims":
			// Add post here later
			Route{Methods:Methods{"GET"}, Handler: withJson(withCQRS(Handler(c.HandleDims)))},
	}
}

// GameController is the controller for a game of life board.
type GameController struct {
	game          *gol.Board
	cqrs          cqrs
	routes        Routes
}


func (c *GameController) HandleDims(ctx Context) Response {
	// No mapping for other methods as of yet
	if ctx.Request.Method == "GET" {
		if err := json.NewEncoder(ctx.Response).Encode(c.game.Dims); err != nil {
			return InternalServerError()
		}
	}
	return StatusNoContent()
}

// HandleBoardClear clears the board of any alive points
func (c *GameController) HandleBoardClear(ctx Context) Response {
	c.game.Clear()
	return StatusNoContent()
}

// HandleBoardCreate creates structures on the board at a given point.
func (c *GameController) HandleBoardCreate(ctx Context) Response {
	log.Println(ctx.Response.Header())
	r := &gameRequest{}
	if err := json.NewDecoder(ctx.Request.Body).Decode(r); err != nil {
		log.Println(err)
		return BadRequest(err)
	}
	if err := c.game.Handle(r); err != nil {
		log.Println(err)
		return BadRequest(err)
	}
	return StatusNoContent()
}

type gameRequest struct {
	Pattern gol.GolPattern `json:"pattern"`
	Point   image.Point    `json:"point"`
}

func (g gameRequest) At() image.Point {
	return g.Point
}

func (g gameRequest) Type() gol.GolPattern {
	return g.Pattern
}

