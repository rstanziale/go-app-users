package handlers

import (
	adapters "app-users/adapters/users"
	"app-users/enums"
	"app-users/model"
	"encoding/json"
	"log"
	"net/http"
)

func ApiCreate(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	userId, err := adapters.CreateUser(name, email, password)
	if err != nil {
		log.Println(err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ResponseError(enums.IMPOSSIBLE_SIGN_UP))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.ResponseMessage(userId))
}