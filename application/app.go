package application

import (
	"company/bab/utils/sms"
	"echo-session"
	"my/gamechanger/src/user"

	"fmt"
	"log"

	"github.com/echo-contrib/pongor"
	"github.com/jeyem/mogo"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	BP           = echo.New()
	DB           = dbConnection()
	sessionStore = session.NewCookieStore([]byte(SecureKey))
	Middlewares  = []echo.MiddlewareFunc{
		session.Sessions(AppCookieName, sessionStore),
		middleware.Logger(),
		// middleware.Recover(),
	}
	Sms = sms.New(SMSAPIKey)
)

func Run() {
	template()
	static()
	BP.Use(Middlewares...)
	apps()
	BP.Logger.Fatal(BP.Start(fmt.Sprintf(":%d", Config.Port)))
}

func dbConnection() *mogo.DB {
	uri := fmt.Sprintf("%s:%d/%s", Config.MongoHost, Config.MongoPort,
		Config.MongoDB)
	db, err := mogo.Conn(uri)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func template() {
	r := pongor.GetRenderer(pongor.PongorOption{

		Reload: true,

		Directory: TemplatePath,
	})
	BP.Renderer = r
}

func static() {
	BP.Static("/static", StaticPath)
}

func apps() {
	user.Load(BP, DB)
}
