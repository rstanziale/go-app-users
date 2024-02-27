package handlers

import (
	adapters "app-users/adapters/users"
	"app-users/enums"
	"app-users/managers"
	"html/template"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, _ := template.ParseFiles("templates/layout.html", "templates/login.html")
		t.Execute(w, nil)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	var user, err = adapters.SearchUser(email, password)
	if err != nil {
		log.Println(err)
		http.Error(w, enums.NOT_FOUND, http.StatusInternalServerError)
		return
	}

	sessionId := uuid.New().String()

	http.SetCookie(w, &http.Cookie{
		Name:  enums.SESSION_ID,
		Value: sessionId,
	})
	managers.SetSession(sessionId, user.Id)

	http.Redirect(w, r, "/users/info", http.StatusSeeOther)
}
