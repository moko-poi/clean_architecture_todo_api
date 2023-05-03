package main

import (
	"fmt"

	"github.com/moko-poi/go-todo-api/db"
	"github.com/moko-poi/go-todo-api/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
