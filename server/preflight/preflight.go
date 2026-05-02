// Package preflight handles cors, etc.
package preflight

import (
	"net/http"
)

func EnableCors(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Origin") == "http://localhost:3000" {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	}
}

func PreflightPipe(w http.ResponseWriter, r *http.Request) {
	EnableCors(w, r)
}
