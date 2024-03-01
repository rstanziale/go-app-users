package routes

import (
	"app-users/model"
	routes "app-users/routes/users"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Initialize routes
func InitRoutes(r *mux.Router) {
	var routers = initRouters()

	for _, route := range routers {
		initRoute(route, r)
	}

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome in GO app users!")
	})
}

// Create a list of Router items
func initRouters() []model.Router {
	var routers = []model.Router{
		routes.NewUserRoutes(),
	}
	
	return routers
}

// Call Router interface method for a router
func initRoute(router model.Router, r *mux.Router) {
	router.InitRoutes(r)
}