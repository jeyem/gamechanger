package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
)

func authPage(e echo.Context) error {
	c := e.(*Context)
	return c.Render(200, "", "")
}

func login(e echo.Context) error {
	c := e.(*Context)
	form := new(loginForm)
	if err := c.Bind(form); err != nil {
		return c.String(403, "bind error")
	}
	if _, err := govalidator.ValidateStruct(form); err != nil {
		return c.String(403, "validate error")
	}
	user := new(Model)
	if err := user.Auth(form.Email, form.Password); err != nil {
		return c.String(403, "auth error")
	}
	c.session.Set(AuthCookieName, user.ID.Hex())
	c.session.Save()
	next := c.nextURL()
	if next != "" {
		return c.Redirect(302, next)
	}
	return c.String(200, "ok")
}

func register(e echo.Context) error {
	c := e.(*Context)
	form := new(registerForm)
	if err := c.Bind(form); err != nil {
		return c.String(403, err.Error())
	}
	if err := form.Validate(); err != nil {
		return c.String(403, err.Error())
	}
	user := new(Model)
	form.initModel(user)
	if err := user.Save(); err != nil {
		return c.String(403, err.Error())
	}
	c.session.Set(AuthCookieName, user.ID.Hex())
	c.session.Save()
	return c.String(200, "success")

}

func logout(e echo.Context) error {
	c := e.(*Context)
	c.session.Delete(AuthCookieName)
	c.session.Save()
	return c.String(200, "OK")
}
