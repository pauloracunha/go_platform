package routes

import (
	"net/http"
	"portal/internal/handlers"

	"github.com/gorilla/mux"
)

func RegisterUsersRoutes(router *mux.Router) {
	router.HandleFunc("/users", handlers.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/users", handlers.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", handlers.GetUserByID).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods(http.MethodDelete)
}
