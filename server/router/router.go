// Package router
package router

import (
	"net/http"
	"strings"
)

func GetReqMainRoute(r *http.Request) string {
	q := r.URL.String()
	s := strings.Split(q, "/")

	// INFO: http://localhost:3000/ is "len(s) == 2"

	if len(s) > 3 {
		return "invalid-url"
	}

	return s[1]
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
