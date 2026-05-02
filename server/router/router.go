// Package router
package router

import (
	"net/http"
	"slices"
	"strings"
)

func getMainRoute(r *http.Request) string {
	q := r.URL.String()

	s := strings.Split(q, "/")

	// the > 3 is to prevent e.g. "login/xxx/xxx"... from working
	// however currently, "login/[garbage]" will still work
	// but not "login/[garbage]/[garbage]"
	// INFO: localhost:3000 defaults for "len(s) == 2"
	// TODO: implement query params to improve this
	if len(s) == 2 && q == "/" {
		return "/"
	}

	if len(s) > 3 {
		return "invalid-url"
	}

	mainRoute := s[1]
	return mainRoute
}

func RouterPipe(r *http.Request) (bool, string) {
	routes := []string{
		"home",
		"login",
		"product",
		"interaction",
		// "service" does not serve any page, see page.go for implementation
		"service",
	}

	mainRoute := getMainRoute(r)

	if mainRoute == "/" {
		return false, "/"
	} else if slices.Contains(routes, mainRoute) {
		return true, mainRoute
	} else {
		return false, ""
	}

}
