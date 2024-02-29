package handlers

import (
	adapters "app-users/adapters/users"
	"app-users/enums"
	"app-users/model"
	"encoding/json"
	"log"
	"net/http"
)

func ApiGetUsers(w http.ResponseWriter, r *http.Request) {
	adapter := adapters.NewUserAdapter()

	users, err := adapter.GetUsers()
	if err != nil {
		log.Println(err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ResponseError(enums.USER_NOT_FOUND))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.ResponseResult(users))
}
