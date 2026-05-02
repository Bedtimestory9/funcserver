package service

import (
	"fmt"
	"funcserver/server/db"
	"funcserver/server/session"
	"io"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func loginValidation(mux *http.ServeMux, conn *pgx.Conn, s *session.SessionManager) {
	mux.HandleFunc("POST /service/validation/login-validation", func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("No body has been posted.")
		}

		var records []db.UserRecord
		db.DecodeJSON(b, &records)
		fmt.Println(records[0])

		u := records[0].Username
		p := records[0].Password
		err = db.QueryUser(conn, u, p)
		if err != nil {
			w.Write([]byte("Username or password incorrect"))
		} else {
			w.Write([]byte("Log in successfully"))
			s.AddUserToSession(u)
		}
	})
}
