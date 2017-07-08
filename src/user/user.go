package user

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

const (
	AuthCookieName = "user"
	NextUrlCookie  = "next_url"
)

func routes(e *echo.Echo) {
	g := application.BP.Group("/user/")
	g.GET("auth", authPage, LoginRequired).Name = "auth"
	g.POST("auth/login", login, NotLoginRequired).Name = "login"
	g.POST("auth/register", register, NotLoginRequired).Name = "register"
	g.GET("auth/logout", logout, LoginRequired).Name = "logout"
}

func Load(e *echo.Echo, d *mogo.DB) {
	db = d
	e.Use(initContext)
	routes(e)
}
