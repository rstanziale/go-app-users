package routes

import (
	"app-users/handlers/users"
	"net/http"

	"github.com/gorilla/mux"
)

const BASE_USER_ROUTE = "users"

func InitUsersRoutes(r *mux.Router) {
	s := r.PathPrefix("/" + BASE_USER_ROUTE).Subrouter()

	s.HandleFunc("/login", handlers.Login).Methods(http.MethodGet, http.MethodPost)
	s.HandleFunc("/logout", handlers.Logout).Methods(http.MethodPost)
	s.HandleFunc("/info", handlers.Info).Methods(http.MethodGet, http.MethodPost)
	s.HandleFunc("/register", handlers.Register).Methods(http.MethodGet, http.MethodPost)
}
