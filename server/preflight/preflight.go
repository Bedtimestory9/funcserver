// Package preflight handles cors, etc.
package preflight

import (
	"fmt"
	"net/http"
)

func EnableCors(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header)

	if r.Header.Get("Origin") == "http://localhost:3000" {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	}
}

func PreflightPipe(w http.ResponseWriter, r *http.Request) {
	EnableCors(w, r)
}
