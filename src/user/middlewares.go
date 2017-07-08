package user

import "github.com/labstack/echo"

func LoginRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		c := e.(*Context)
		if !c.authenticate {
			path := c.Request().URL.Path
			c.session.Set(NextUrlCookie, path)
			c.session.Save()
			return c.Redirect(302, "/user/auth")
		}
		return next(c)
	}
}

func NotLoginRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		c := e.(*Context)
		if c.authenticate {
			return c.Redirect(302, "/")
		}
		return next(c)
	}
}
