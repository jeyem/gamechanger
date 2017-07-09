package game

import (
	"my/gamechanger/src/user"
	"strconv"

	"github.com/labstack/echo"
)

func list(e echo.Context) error {
	params := e.QueryParam("q")
	page, _ := strconv.Atoi(e.QueryParam("page"))
	if page == 0 {
		page = 1
	}
	games := LoadGames(10, page, params)
	return e.Render(200, "games.html", echo.Map{
		"games": games,
	})
}

func viewGame(e echo.Context) error {
	gameID := e.Param("game_id")
	if !IsObjectIdHex(gameID) {
		return e.Render(403, "403.html", echo.Map{})
	}
	game := new(Model)
	if err := game.Load(gameID); err != nil {
		return e.Render(404, "404.html", echo.Map{})
	}
	return e.Render(200, "game-view.html", echo.Map{
		"game": game,
	})
}

func haveGame(e echo.Context) error {
	gameID := e.Param("game_id")
	if !IsObjectIdHex(gameID) {
		return e.Render(403, "403.html", echo.Map{})
	}
	game := new(Model)
	if err := game.Load(gameID); err != nil {
		return e.Render(404, "404.html", echo.Map{})
	}
	c := e.(*user.Context)
	geme.AppendOwner(c.User.ID)
	if err := game.Save(); err != nil {
		return e.Render(403, "game-view.html", echo.Map{
			"error": err.Error(),
			"game":  game,
		})
	}
	return e.Render(200, "game-view.html", echo.Map{
		"game": game,
		"msg":  "success append",
	})
}
