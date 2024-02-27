package handlers

import (
	adapters "app-users/adapters/users"
	"app-users/enums"
	"app-users/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ApiGetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		log.Println(err)
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ResponseError(enums.USER_NOT_FOUND))

		return
	}

	var user model.User
	user, err = adapters.SearchUserById(userIdInt)
	if err != nil {
		log.Println(err)
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ResponseError(enums.USER_NOT_FOUND))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.ResponseMessage(user))
}
