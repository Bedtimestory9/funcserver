// Package page serve [layout] + [views] template
package page

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

// type TMPLData struct {
// 	Title    string
// 	UserID   string
// 	IsAuthed bool
// }

type MainRoute struct {
	Route string
}

// func PassIfAuthTMPLData(authed bool) TMPLData {
// 	var tmplData TMPLData
//
// 	if authed {
// 		tmplData = TMPLData{
// 			Title:    "User Session",
// 			UserID:   "User logged in",
// 			IsAuthed: true,
// 		}
// 	} else {
// 		tmplData = TMPLData{
// 			Title:    "Guest Session",
// 			UserID:   "Please Log In",
// 			IsAuthed: false,
// 		}
// 	}
//
// 	return tmplData
// }

func serveTemplate(w http.ResponseWriter, pageRoute string) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	routeParam := strings.Split(pageRoute, "/")[1]

	// the order of the files matter, base first
	tmpl, err := template.ParseFiles(wd+"/server/views/layout.html", wd+"/server/views/"+routeParam+"/"+routeParam+".html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func (m MainRoute) pageHandler(w http.ResponseWriter, r *http.Request) {
	switch m.Route {
	case "/":
		m.Route = "home"
		serveTemplate(w, m.Route)
	case "/service":
		w.WriteHeader(404)
		w.Write([]byte("Page not found"))
	default:
		serveTemplate(w, m.Route)
	}
}

func RouterPipe(mux *http.ServeMux) {
	routes := []string{
		"/",
		"/home",
		"/login",
		"/product",
		"/interaction",
		"/signup",
		// "service" does not serve any page
		"/service",
	}

	mainRoute := MainRoute{}

	for _, r := range routes {
		mainRoute.Route = r
		mux.HandleFunc(r, mainRoute.pageHandler)
	}
}
