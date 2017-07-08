package game

import (
	"strconv"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
)

func list(e echo.Context) error {
	params := e.QueryParam("q")
	page, _ := strconv.Atoi(e.QueryParam("page"))
	if page == 0 {
		page = 1
	}
	games := LoadGames(10, page, params)
	return e.Render(200, "games.html", pongo2.Context{
		"games": games,
	})
}
