package models

import (
	"errors"
	"goTimisoaraBackend/db"
	"golang.org/x/crypto/bcrypt"
	"html"
	"log"
	"strings"
	"time"
)

type User struct {
	UserId    uint32    `gorm:"primary_key;auto_increment" json:"userId"`
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	FullName  string    `gorm:"size:100; not null" json:"fullName"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (user *User) VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (user *User) BeforeSave() error {
	log.Println(user.Password)
	hashedPassword, err := Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func (user *User) Prepare() {
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	user.FullName = html.EscapeString(strings.TrimSpace(user.FullName))
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
}

func (user *User) SaveUser() (*User, error) {
	var err error

	database := db.GetDB()
	err = database.Debug().Create(&user).Error

	if err != nil {
		return &User{}, err
	}

	return user, nil
}

func (user *User) FindAllUsers() (*[]User, error) {
	var err error
	var users []User
	database := db.GetDB()

	err = database.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}

	return &users, err
}

func (user *User) FindUserById(userId uint32) (User, error) {
	var err error
	var userResult User
	database := db.GetDB()

	err = database.Debug().Model(&User{}).Where("user_id = ?", userId).Limit(1).Find(&userResult).Error

	if err != nil {
		return User{}, err
	}

	return userResult, err
}

func (user *User) FindUserByUsername(username string) (User, error) {
	var err error
	var result User

	database := db.GetDB()

	err = database.Debug().Where("username = ?", username).First(&result).Error

	if err != nil {
		return User{}, errors.New("user not found")
	}

	return result, err
}

func (user *User) UpdateUser() (*User, error) {
	database := db.GetDB()

	databaseResult := database.Debug().Model(&User{}).Where("user_id = ?", user.UserId).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"username":   user.Username,
			"email":      user.Email,
			"full_name":  user.FullName,
			"updated_at": time.Now(),
		},
	)

	if databaseResult.Error != nil {
		return &User{}, databaseResult.Error
	}

	return user, nil
}

func (user *User) DeleteUserById(id uint32) error {
	database := db.GetDB()
	databaseResult := database.Debug().Model(&User{}).Where("user_id = ?", id).Take(&User{}).Delete(&User{})

	if databaseResult.Error != nil {
		return databaseResult.Error
	}

	return nil
}
