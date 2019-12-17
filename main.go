package main

import (
	"flag"
	"fmt"
	"goTimisoaraBackend/config"
	"goTimisoaraBackend/db"
	"goTimisoaraBackend/models"
	"goTimisoaraBackend/server"
	"os"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}

	flag.Parse()
	config.Init(*environment)
	db.Init()
	db.GetDB().AutoMigrate(&models.User{}, &models.Article{}, &models.Tag{}, &models.Author{})
	server.Init()
}
