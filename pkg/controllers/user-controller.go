package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Sambit99/Basic-RestAPI-Go-MongoDB/pkg/models"
	"github.com/Sambit99/Basic-RestAPI-Go-MongoDB/pkg/utils"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/v2/bson"
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
func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user = &models.User{}

	utils.ParseBody(r, user)
	id := user.CreateUser()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId := params["userId"]

	isDeleted := models.DeleteUser(userId)

	if !isDeleted {
		log.Fatal("Something went wrong while deleting record")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Record deleted successfully")
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	userId := params["userId"]

	var user models.User

	err := utils.ParseBody(r, &user)

	if err != nil {
		log.Fatal("Error while parsing body")
	}

	id, err := bson.ObjectIDFromHex(userId)

	if err != nil {
		log.Fatal("Error while converting user Id")
	}

	user.ID = id

	isUpdated := models.UpdateUser(&user)

	if !isUpdated {
		log.Fatal("Error while updating user")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("User updated successfully")
}
