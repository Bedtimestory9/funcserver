// Package server handles server
package server

import (
	"funcserver/server/misc"
	"funcserver/server/page"
	"funcserver/server/service"
	"net/http"
)

func NewHTTPServer(addr string) *http.Server {
	p := page.NewPage()

	r := http.NewServeMux()

	r.Handle("/styles.css", http.FileServer(http.Dir(misc.GetCmdDir()+"/public")))
	r.Handle("/scripts/", http.FileServer(http.Dir(misc.GetCmdDir()+"/public")))

	// r.HandleFunc("/", service.SetCookieHandler)
	r.HandleFunc("/home", p.HomePageHandler)
	r.HandleFunc("/login", p.LoginPageHandler)
	r.HandleFunc("POST /login", p.PostLoginHandler)
	r.HandleFunc("/signup", p.SignupPageHandler)
	r.HandleFunc("/interaction", p.InteractionPageHandler)

	s := service.NewService()

	r.HandleFunc("POST /service/login", s.LoginServiceHandler)
	r.HandleFunc("POST /service/signup", s.SignupServiceHandler)
	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}
