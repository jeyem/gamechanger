package game

import (
	"company/bab/application"
	session "echo-session"

	"github.com/jeyem/mogo"
	"github.com/labstack/echo"
)

var (
	db   *mogo.DB
	sess session.CookieStore
)

func routes(e *echo.Echo) {
	g := application.BP.Group("/game/")
	g.GET("list", list)
}

func Load(e *echo.Echo, d *mogo.DB) {
	db = d
	routes(e)
}
