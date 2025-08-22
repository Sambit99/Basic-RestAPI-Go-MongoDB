package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Sambit99/Basic-RestAPI-Go-MongoDB/pkg/models"
	"github.com/gorilla/mux"
)

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUser()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId := params["userId"]

	user := models.GetUserById(userId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
