package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"goTimisoaraBackend/db"
	"golang.org/x/crypto/bcrypt"
	"html"
	"log"
	"strings"
	"time"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	FullName  string    `gorm:"size:100; not null" json:"FullName"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (u *User) VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
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
	case "login":
		if u.Password == "" {
			return errors.New("required Password")
		}
		if u.Email == "" {
			return errors.New("required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid Email")
		}
		return nil

	default:
		if u.Username == "" {
			return errors.New("required Username")
		}
		if u.Password == "" {
			return errors.New("required Password")
		}
		if u.Email == "" {
			return errors.New("required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid Email")
		}

		return nil
	}
}

func (u *User) SaveUser() (*User, error) {

	var err error
	database := db.GetDB()
	err = database.Debug().Create(&u).Error

	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func (u *User) FindAllUsers() (*[]User, error) {
	var err error
	var users []User
	database := db.GetDB()

	err = database.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}

	return &users, err
}

func (u *User) FindUserById(uid string) (*User, error) {
	var err error
	database := db.GetDB()

	err = database.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error

	if err != nil {
		return &User{}, errors.New("user not found")
	}

	return u, err
}

func (u *User) FindUserByUsername(username string) (*User, error) {
	var err error
	database := db.GetDB()

	err = database.Debug().Model(User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return &User{}, errors.New("user not found")
	}

	return u, err
}

func (u *User) UpdateAUser(uid uint32) (*User, error) {
	database := db.GetDB()

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}

	databaseResult := database.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":  u.Password,
			"Username":  u.Username,
			"email":     u.Email,
			"update_at": time.Now(),
		},
	)

	if databaseResult.Error != nil {
		return &User{}, databaseResult.Error
	}
	// This is the display the updated user
	err = databaseResult.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error

	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) DeleteAUser(uid uint32) (int64, error) {
	database := db.GetDB()
	databaseResult := database.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if databaseResult.Error != nil {
		return 0, databaseResult.Error
	}
	return databaseResult.RowsAffected, nil
}
