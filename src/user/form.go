package user

import (
	"my/gamechanger/utils/auth"

	"github.com/asaskevich/govalidator"
)

type loginForm struct {
	Email    string `valid:"email,required" json:"email" form:"email"`
	Password string `valid:"required" json:"password" form:"password"`
}

type registerForm struct {
	Email    string `valid:"email,required" json:"email" form:"email"`
	Password string `valid:"required" json:"password" form:"password"`
	Fullname string `json:"fullname" form:"fullname"`
}

func (f *registerForm) Validate() error {
	if _, err := govalidator.ValidateStruct(f); err != nil {
		return err
	}
	return nil
}

func (f *registerForm) initModel(model *Model) {
	model.Fullname = f.Fullname
	model.Password = auth.MakePassword(f.Password)
	model.OrginEmail = f.Email
	model.Email = auth.EmailFixer(f.Email)

}
