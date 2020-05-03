package config

import (
	"database/sql"
	"net/http"

	"github.com/felipehfs/clean-api/presenters"
	"github.com/felipehfs/clean-api/repositories/pg"
	"github.com/felipehfs/clean-api/usecases"
	"github.com/gorilla/mux"
)

// Server operates the server
type Server struct {
	DB     *sql.DB
	Router *mux.Router
}

// Run initializes the server and the port
func (s Server) Run(port string) {
	userRepository := pg.UserRepository{
		DB: s.DB,
	}

	userUsecase := usecases.UserService{
		Repository: userRepository,
	}

	userHandler := presenters.UserHandler{
		Service: userUsecase,
	}

	s.Router.HandleFunc("/api/register", userHandler.Register).Methods("POST")
	http.ListenAndServe(port, s.Router)
}
