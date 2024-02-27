package routes

import (
	routes "app-users/routes/users"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes(r *mux.Router) {
	routes.InitUsersRoutes(r)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
}
