// Package service deals with specific
// functionalities contextually
// eg. Validation .etc
package service

import (
	"net/http"

	"github.com/jackc/pgx/v5"
)

func ServicePipe(mux *http.ServeMux, conn *pgx.Conn) {
	loginValidation(mux, conn)
	signupValidation(mux, conn)
	getUserMood(mux, conn)
}
