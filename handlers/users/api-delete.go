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

func ApiDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		log.Println(err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ResponseError(enums.USER_NOT_DELETED))
	}

	adapter := adapters.NewUserAdapter()

	err = adapter.DeleteUser(userIdInt)
	if err != nil {
		log.Println(err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ResponseError(enums.USER_NOT_DELETED))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.ResponseOK())
}
