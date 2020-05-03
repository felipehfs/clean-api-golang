package main

import (
	"github.com/felipehfs/clean-api/config"
	"github.com/gorilla/mux"
)

func main() {
	db, err := config.GetPostgresInstance()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	r := mux.NewRouter()
	server := config.Server{
		DB:     db,
		Router: r,
	}

	server.Run(":8080")
}
