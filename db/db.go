package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goTimisoaraBackend/config"
	"log"
)

var db *gorm.DB

func Init() {
	c := config.GetConfig()

	databaseCredentials := fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local", c.GetString("db.username"), c.GetString("db.password"), c.GetString("db.hostname"), c.GetString("db.database"))
	dbConnection, err := gorm.Open("mysql", databaseCredentials)

	if err != nil {
		log.Println("Faulty connection")
	}

	db = dbConnection
}

func GetDB() *gorm.DB {
	return db
}
