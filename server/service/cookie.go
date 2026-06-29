package service

import "net/http"

func SetCookieHandler(w http.ResponseWriter, r http.Request) {
	cookie := http.Cookie{
		Name:     "functionalServerCookie",
		Value:    "functionalCookie",
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
	}

	http.SetCookie(w, &cookie)
}
