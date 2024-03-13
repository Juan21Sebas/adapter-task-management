package main

import (
	"adapter-task-management/cmd/config/db"
	"adapter-task-management/internal/adapters/handlers/http"
	"log"
)

func main() {
	dbInstance, err := db.NewSQLiteDB()
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	http.RunServer(dbInstance)
}
