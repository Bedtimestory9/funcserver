// Package page serve [layout] + [views] template
package page

import (
	"funcserver/server/router"
	"funcserver/server/session"
	"html/template"
	"log"
	"net/http"
	"os"
)

func nestPageInLayout(w http.ResponseWriter, route string, tmplData session.TMPLData) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// the order of the files matter, base first
	tmpl, err := template.ParseFiles(wd+"/server/views/layout.html", wd+"/server/views/"+route+"/"+route+".html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, tmplData)
	if err != nil {
		panic(err)
	}
}

func serveRouterPage(w http.ResponseWriter, routeExist bool, route string, tmplData session.TMPLData) {
	// "service" is skipped from serving page
	if routeExist && route != "service" {
		nestPageInLayout(w, route, tmplData)
	} else if route == "/" {
		route = "home"
		nestPageInLayout(w, route, tmplData)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("Page Not Found"))
	}
}

func PagePipe(mux *http.ServeMux, s *session.SessionManager) {
	mux.Handle("/styles.css", http.FileServer(http.Dir("public")))
	mux.Handle("/scripts/", http.FileServer(http.Dir("public")))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// It doesn't do anything for now,
		// since EnableCors() does not work
		// on other service routes, and "/" does not need it
		// preflight.PreflightPipe(w, r)

		tmplData := session.SessionPipe(w, r, s)

		re, route := router.RouterPipe(r)

		serveRouterPage(w, re, route, tmplData)
	})
}
