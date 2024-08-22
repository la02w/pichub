package main

import (
	"pichub/model"
	"pichub/router"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	model.InitDb()
	router.InitRouter()
}
