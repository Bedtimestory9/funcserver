package service

import (
	"encoding/json"
	"fmt"
	"funcserver/server/db"
	"funcserver/server/session"
	"io"
	"net/http"
	"net/mail"

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

		u := records[0].Username
		p := records[0].Password

		res, err := db.QueryLoginUser(conn, u, p)
		if err != nil {
			fmt.Println("error querying login user")
		}

		jsonData, err := json.Marshal(res)
		if err != nil {
			fmt.Println("failed marshalling data to json")
		}

		if err != nil {
			w.WriteHeader(401)
		} else {
			w.WriteHeader(303)
			s.AddUserToSession(u)
		}

		w.Write(jsonData)
	})
}

func signupValidation(mux *http.ServeMux, conn *pgx.Conn, s *session.SessionManager) {
	mux.HandleFunc("POST /service/validation/signup-validation", func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("No body has been posted.")
		}

		var records []db.UserSignupRecord
		db.DecodeJSON(b, &records)

		e := records[0].Email
		u := records[0].Username
		p := records[0].Password
		a := records[0].Age

		res := db.UserSignupResponse{}

		if a < 18 {
			res.AgeError = "age must be equal and greater than 18"

			w.WriteHeader(401)
			jsonData, err := json.Marshal(res)
			if err != nil {
				fmt.Println("failed marshalling data to json")
			}
			w.Write(jsonData)

			return
		}

		_, err = mail.ParseAddress(e)

		if err != nil {
			res.EmailError = "Must be an email address"

			w.WriteHeader(401)
			jsonData, err := json.Marshal(res)
			if err != nil {
				fmt.Println("failed marshalling data to json")
			}
			w.Write(jsonData)

			return
		}

		res, shouldContinue := db.QuerySignupUser(conn, e, u)

		if shouldContinue {
			db.InsertSignupUser(conn, e, u, p, a, &res)

			w.WriteHeader(303)
			jsonData, err := json.Marshal(res)
			if err != nil {
				fmt.Println("failed marshalling data to json")
			}
			s.AddUserToSession(u)
			w.Write(jsonData)
		} else {
			w.WriteHeader(401)
			jsonData, err := json.Marshal(res)
			if err != nil {
				fmt.Println("failed marshalling data to json")
			}
			w.Write(jsonData)
		}

	})
}
