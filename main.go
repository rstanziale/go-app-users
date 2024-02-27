package main

import (
	"app-users/middlewares"
	"app-users/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	r := mux.NewRouter()

	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(middlewares.Logging)

	routes.InitRoutes(r)

	log.Fatal(http.ListenAndServe(":9090", r))
}
