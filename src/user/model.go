package user

import (
	"errors"
	"my/gamechanger/utils/auth"
	"my/gamechanger/utils/random"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/labstack/echo"
)

var (
	ErrorTokenExpired = errors.New("Token has been expired")
)

type Model struct {
	ID                  bson.ObjectId `bson:"_id,omitempty"`
	CreatedAt           time.Time     `bson:"created_at"`
	AuthToken           string        `bson:"auth_token"`
	AuthTokenExpireTime time.Time     `bson:"auth_token_time"`
	Email               string        `bson:"email"`
	OrginEmail          string        `bson:"orgin_email"`
	Password            string        `bson:"password"`
	ForgetPasswordKey   string        `bson:"forget_password_key"`
	Fullname            string        `bson:"fullname"`
	EmailVerifyKey      string        `bson:"mail_verify_key"`
	Marital             string        `bson:"marital"`
	Website             string        `bson:"website"`
	Cellphone           string        `bson:"cellphone"`
	LastLogin           time.Time     `bson:"last_login"`
	Keywords            []string      `bson:"keywords"`
	IsActive            bool          `bson:"is_active"`
}

var (
	errorModelOrPass    = errors.New("email or password not matched")
	errorModelNotActive = errors.New("user is not active")
	errorDuplicateModel = errors.New("user already exists")
)

func (Model) Meta() []mgo.Index {
	return []mgo.Index{
		{
			Key:    []string{"email"},
			Unique: true,
		},
	}
}

func (c *Model) Auth(email, password string) error {
	email = auth.EmailFixer(email)
	if err := c.LoadWithMail(email); err != nil {
		return errorModelOrPass
	}
	if ok := auth.CheckPassword(password, c.Password); !ok {
		return errorModelOrPass
	}
	if !c.IsActive {
		return errorModelNotActive
	}
	c.LastLogin = time.Now()
	return c.Update()
}

func (c *Model) LoadWithMail(email string) error {
	return db.Where(bson.M{
		"email": email,
	}).Find(c)
}

func (c *Model) LoadWithCellphone(cellphone string) error {
	return db.Where(bson.M{
		"cellphone": cellphone,
	}).Find(c)
}

func (c *Model) Save() error {
	if c.ID.Valid() {
		return c.Update()
	}
	if err := c.checkDuplicate(); err != nil {
		return err
	}
	c.CreatedAt = time.Now()
	c.setKeywords()
	c.IsActive = true
	return db.Create(c)
}

func (c Model) checkDuplicate() error {
	user := new(Model)
	if c.Email == "" && c.Cellphone == "" {
		return errors.New("no primary key for user")
	}
	if c.Email != "" {
		if err := user.LoadWithMail(c.Email); err == nil {
			return errorDuplicateModel
		}
	}
	if c.Cellphone != "" {
		if err := user.LoadWithCellphone(c.Cellphone); err == nil {
			return errorDuplicateModel
		}
	}
	return nil
}

func (c *Model) Update() error {
	return db.Update(c)
}

func (c *Model) Load(id interface{}) error {
	if val, ok := id.(string); ok {
		if !bson.IsObjectIdHex(val) {
			return errors.New("id error")
		}
		id = bson.ObjectIdHex(val)
	}
	return db.Get(c, id)
}

func (c *Model) setKeywords() {
	c.Keywords = []string{
		c.Email,
		c.Fullname,
	}
}

func (c *Model) ShowableName() string {
	if c.Fullname != "" {
		return c.Fullname
	}
	return c.Email
}

func (c *Model) Rest() echo.Map {
	resp := echo.Map{
		"showable_name": c.ShowableName(),
		"email":         c.Email,
		"id":            c.ID,
	}

	if c.Cellphone != "" {
		resp["cellphone"] = c.Cellphone
	}

	return resp
}

func (c *Model) GenerateAuthKey() error {
	c.AuthTokenExpireTime = time.Now().Add(time.Minute * 5)
	c.AuthToken = random.Rand(6)
	return c.Save()
}

func (c *Model) AuthWithKey(cellphone, token string) error {
	if err := db.Where(bson.M{
		"cellphone":  cellphone,
		"auth_token": token,
	}).Find(c); err != nil {
		return err
	}
	if time.Now().After(c.AuthTokenExpireTime) {
		return ErrorTokenExpired
	}
	c.AuthToken = ""
	return c.Save()

}

func (c *Model) RequestAuthKey(cellphone string) error {
	if err := db.Where(bson.M{
		"cellphone": cellphone,
	}).Find(c); err != nil {
		c.Cellphone = cellphone
	}
	return c.GenerateAuthKey()
}
