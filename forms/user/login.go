package userForms

import (
	"errors"
)

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserLogin) Validate() error {
	if u.Username == "" {
		return errors.New("required Username")
	}
	if u.Password == "" {
		return errors.New("required Password")
	}

	return nil
}
