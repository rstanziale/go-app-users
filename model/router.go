package model

import "github.com/gorilla/mux"

type Router interface {
	InitRoutes(r *mux.Router)
}
