package userForms

import (
	"errors"
	"github.com/badoux/checkmail"
)

type UserRegister struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
	Password string `json:"password"`
}

func (u *UserRegister) Validate() error {
	if u.Username == "" {
		return errors.New("required Username")
	}

	if u.Password == "" {
		return errors.New("required Password")
	}

	if u.Email == "" {
		return errors.New("required Email")
	}

	if u.FullName == "" {
		return errors.New("required Full Name")
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("invalid Email")
	}

	return nil
}
