package user

import (
	session "echo-session"

	"github.com/labstack/echo"
)

type Context struct {
	echo.Context
	authenticate bool
	User         Model
	session      session.Session
}

func (c *Context) loadUser() {
	val := c.session.Get(AuthCookieName)
	id, ok := val.(string)
	if !ok || id == "" {
		return
	}
	if err := c.User.Load(id); err == nil {
		c.authenticate = true
	}
}

func (c *Context) nextURL() string {
	val := c.session.Get(NextUrlCookie)
	url, ok := val.(string)
	if !ok || url == "" {
		return ""
	}
	c.session.Delete(NextUrlCookie)
	c.session.Save()
	return url
}

func initContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		c := &Context{Context: e}
		c.session = session.Default(c)
		c.loadUser()
		return next(e)
	}
}
