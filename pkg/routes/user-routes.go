package routes

import (
	"github.com/Sambit99/Basic-RestAPI-Go-MongoDB/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(r *mux.Router) {
	r.HandleFunc("/user", controllers.GetAllUser).Methods("GET")
	r.HandleFunc("/user/{userId}", controllers.GetUserById).Methods("GET")
	r.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/user", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{userId}", controllers.DeleteUser).Methods("DELETE")
}
