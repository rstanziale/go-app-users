package routes

import (
	handlers "app-users/handlers/users"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type UserRoutes struct {
	baseUrl string
	apiUrl string
}

func NewUserRoutes() *UserRoutes {
	userRoutes := UserRoutes{"users", "api"}
	return &userRoutes
}

func (userRoutes *UserRoutes) InitRoutes(r *mux.Router) {
	var baseUserUrl = fmt.Sprintf("/%s", userRoutes.baseUrl)
	s := r.PathPrefix(baseUserUrl).Subrouter()

	s.HandleFunc("/login", handlers.Login).Methods(http.MethodGet, http.MethodPost)
	s.HandleFunc("/logout", handlers.Logout).Methods(http.MethodPost)
	s.HandleFunc("/info", handlers.Info).Methods(http.MethodGet, http.MethodPost)
	s.HandleFunc("/register", handlers.Register).Methods(http.MethodGet, http.MethodPost)

	var apiUserUrl = fmt.Sprintf("/%s/%s", userRoutes.apiUrl, userRoutes.baseUrl)
	sApi := r.PathPrefix(apiUserUrl).Subrouter()

	sApi.HandleFunc("/get", handlers.ApiGetUsers).Methods(http.MethodGet)
	sApi.HandleFunc("/get/{id}", handlers.ApiGetUser).Methods(http.MethodGet)
	sApi.HandleFunc("/create", handlers.ApiCreate).Methods(http.MethodPost)
	sApi.HandleFunc("/update", handlers.ApiUpdate).Methods(http.MethodPut)
	sApi.HandleFunc("/delete/{id}", handlers.ApiDelete).Methods(http.MethodDelete)
}
