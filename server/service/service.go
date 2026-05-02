// Package service deals with specific
// functionalities contextually
// eg. Validation .etc
package service

import (
	"funcserver/server/session"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func ServicePipe(mux *http.ServeMux, conn *pgx.Conn, s *session.SessionManager) {
	loginValidation(mux, conn, s)
	getUserMood(mux, conn)
}
