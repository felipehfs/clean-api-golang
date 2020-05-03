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

// Create inserts the new data
func (handler UserHandler) Create(w http.ResponseWriter, r *http.Request) {
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

	response := make(map[string]interface{})
	response["data"] = user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
