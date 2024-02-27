package handlers

import (
	adapters "app-users/adapters/users"
	"app-users/enums"
	"app-users/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func ApiUpdate(w http.ResponseWriter, r *http.Request) {
	userId := r.FormValue("userId")
	newName := r.FormValue("name")
	newPassword := r.FormValue("password")

	userIdInt, err := strconv.Atoi(userId)
    if err != nil {
		log.Println(err)
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ResponseError(enums.USER_NOT_UPDATED))

		return
    }

	err = adapters.UpdateUser(userIdInt, newName, newPassword)
	if err != nil {
		log.Println(err)
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ResponseError(enums.USER_NOT_UPDATED))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.ResponseOK())
}