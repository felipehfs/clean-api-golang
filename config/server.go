package config

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
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

func getToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	bearArray := strings.Split(bearToken, " ")
	if len(bearArray) == 2 {
		return bearArray[1]
	}
	return ""
}

func verifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := getToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_TOKEN_CLEAN_API")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func hasAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("authorization") != "" {
			token, err := verifyToken(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if token.Valid {
				handler(w, r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintln(w, "Not authorized")
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Not authorized")
		}
	}
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

	bookRepository := pg.BookRepository{
		DB: s.DB,
	}

	bookHandler := presenters.NewBookHandler(bookRepository)

	s.Router.HandleFunc("/api/register", userHandler.Register).Methods("POST")
	s.Router.HandleFunc("/api/login", userHandler.Login).Methods("POST")

	s.Router.HandleFunc("/api/books", bookHandler.Create).Methods("POST")
	s.Router.HandleFunc("/api/books", bookHandler.Get).Methods("GET")
	s.Router.HandleFunc("/api/books/{id}", hasAuth(bookHandler.Update)).Methods("PUT")
	s.Router.HandleFunc("/api/books/{id}", hasAuth(bookHandler.Remove)).Methods("DELETE")

	http.ListenAndServe(port, s.Router)
}
