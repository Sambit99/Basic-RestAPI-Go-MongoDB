package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Sambit99/Basic-RestAPI-Go-MongoDB/pkg/models"
)

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUser()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
