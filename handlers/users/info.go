package handlers

import (
	adapters "app-users/adapters/users"
	"app-users/enums"
	"app-users/managers"
	"html/template"
	"log"
	"net/http"
)

func Info(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(enums.SESSION_ID)
	if err != nil {
		http.Error(w, enums.NOT_AUTHENTICATED, http.StatusUnauthorized)
	}

	userId, ok := managers.GetSession(cookie.Value)
	if !ok {
		log.Println(err)
		http.Error(w, enums.SESSION_NOT_VALID, http.StatusUnauthorized)
		return
	}

	switch r.Method {
	case http.MethodGet:
		showUserPage(w, userId)

	case http.MethodPost:
		updateUser(r, w, userId)
	}
}

func showUserPage(w http.ResponseWriter, userId int) {
	user, err := adapters.SearchUserById(userId)
	if err != nil {
		log.Println(err)
		http.Error(w, enums.USER_NOT_FOUND, http.StatusNotFound)
		return
	}

	t, _ := template.ParseFiles("templates/layout.html", "templates/user.html")
	t.Execute(w, user)
}

func updateUser(r *http.Request, w http.ResponseWriter, userId int) {
	newName := r.FormValue("name")
	newPassword := r.FormValue("password")

	err := adapters.UpdateUser(userId, newName, newPassword)
	if err != nil {
		log.Println(err)
		http.Error(w, enums.USER_NOT_UPDATED, http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/users/info", http.StatusSeeOther)
}
