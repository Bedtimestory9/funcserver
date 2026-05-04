// Package page serve [layout] + [views] template
package page

import (
	"html/template"
	"log"
	"net/http"
	"os"
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

func nestPageInLayout(w http.ResponseWriter, route string) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// the order of the files matter, base first
	tmpl, err := template.ParseFiles(wd+"/server/views/layout.html", wd+"/server/views/"+route+"/"+route+".html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}


func (m MainRoute) ServePageHandler(w http.ResponseWriter, r *http.Request) {
	// "service" is skipped from serving page
	if m.Route != "/service" {
		nestPageInLayout(w, m.Route)
	} else if m.Route == "/" {
		m.Route = "/home"
		nestPageInLayout(w, m.Route)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("Page Not Found"))
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
		mux.HandleFunc(r, mainRoute.ServePageHandler)
	}
}
