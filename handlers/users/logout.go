package handlers

import (
	"app-users/enums"
	"app-users/managers"
	"log"
	"net/http"
	"time"
)

// Handler to perform logout
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(enums.SESSION_ID)
	if err != nil {
		log.Println(err)
		http.Error(w, enums.NOT_AUTHENTICATED, http.StatusUnauthorized)
		return
	}

	managers.DeleteSession(cookie.Value)

	http.SetCookie(w, &http.Cookie{
		Name:    enums.SESSION_ID,
		Value:   "",
		Expires: time.Unix(0, 0),
	})

	http.Redirect(w, r, "/users/login", http.StatusSeeOther)
}
