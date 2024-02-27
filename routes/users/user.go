package routes

import (
	handlers "app-users/handlers/users"
	"net/http"

	"github.com/gorilla/mux"
)

const BASE_USER_ROUTE = "users"
const BASE_API_ROUTE = "api"

func InitUsersRoutes(r *mux.Router) {
	s := r.PathPrefix("/" + BASE_USER_ROUTE).Subrouter()

	s.HandleFunc("/login", handlers.Login).Methods(http.MethodGet, http.MethodPost)
	s.HandleFunc("/logout", handlers.Logout).Methods(http.MethodPost)
	s.HandleFunc("/info", handlers.Info).Methods(http.MethodGet, http.MethodPost)
	s.HandleFunc("/register", handlers.Register).Methods(http.MethodGet, http.MethodPost)

	sApi := r.PathPrefix("/" + BASE_API_ROUTE + "/" + BASE_USER_ROUTE).Subrouter()

	sApi.HandleFunc("/get", handlers.ApiGetUsers).Methods(http.MethodGet)
	sApi.HandleFunc("/get/{id}", handlers.ApiGetUser).Methods(http.MethodGet)
	sApi.HandleFunc("/create", handlers.ApiCreate).Methods(http.MethodPost)
	sApi.HandleFunc("/update", handlers.ApiUpdate).Methods(http.MethodPut)
	sApi.HandleFunc("/delete/{id}", handlers.ApiDelete).Methods(http.MethodDelete)
}
