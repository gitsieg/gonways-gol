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
		cqrs: cqrs{headerOptions: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "Content-Type",
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
		Route{Methods: Methods{"POST", "OPTIONS"}, Handler: withCQRS(Handler(c.HandleClear))},
		"/game/create":
		Route{Methods: Methods{"POST", "OPTIONS"}, Handler: withCQRS(Handler(c.HandleCreate))},
		"/game/dims":
		// Add post here later
		Route{Methods: Methods{"GET"}, Handler: withJson(withCQRS(Handler(c.HandleDims)))},
		"/game/options":
		Route{Methods: Methods{"GET"}, Handler: withJson(withCQRS(Handler(c.HandleOptions))),},
	}
}

// GameController is the controller for a game of life board.
type GameController struct {
	game   *gol.Board
	cqrs   cqrs
	routes Routes
}

// HandleOptions returns the games supported structural requests.
func (c *GameController) HandleOptions(ctx Context) Response {
	patterns := map[gol.GolPattern]string{}
	for i := gol.Tumbler; i <= gol.TenCellRow; i++ {
		patterns[i] = i.String()
	}
	resp := struct {
		Patterns map[gol.GolPattern]string `json:"patterns"`
	}{
		Patterns: patterns,
	}
	if err := json.NewEncoder(ctx.Response).Encode(resp); err != nil {
		return InternalServerError()
	}
	return StatusOK()
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

// HandleClear clears the board of any alive points
func (c *GameController) HandleClear(ctx Context) Response {
	c.game.Clear()
	return StatusNoContent()
}

// HandleCreate creates structures on the board at a given point.
func (c *GameController) HandleCreate(ctx Context) Response {
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
