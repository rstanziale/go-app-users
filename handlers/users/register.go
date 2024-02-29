package handlers

import (
	adapters "app-users/adapters/users"
	"app-users/enums"
	"log"
	"net/http"
	"text/template"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, _ := template.ParseFiles("templates/layout.html", "templates/login.html")
		t.Execute(w, nil)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	adapter := adapters.NewUserAdapter()

	_, err := adapter.CreateUser(name, email, password)
	if err != nil {
		log.Println(err)
		http.Error(w, enums.IMPOSSIBLE_SIGN_UP, http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/users/info", http.StatusSeeOther)
}
