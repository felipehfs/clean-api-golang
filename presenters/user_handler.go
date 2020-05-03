package presenters

import (
	"encoding/json"
	"net/http"

	"github.com/felipehfs/clean-api/entities"
	"github.com/felipehfs/clean-api/usecases"
)

// UserHandler represents the controllers of the user
type UserHandler struct {
	Service usecases.UserUsecase
}

// Register inserts the new user
func (handler UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := handler.Service.Create(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.ID = id

	response := make(map[string]interface{})
	response["data"] = user

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
