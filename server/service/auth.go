package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"funcserver/server/db"
	"io"
	"net/http"
	"net/mail"
	"os"

	"github.com/jackc/pgx/v5"
)

type Service struct {
	conn *pgx.Conn
}

func NewService() *Service {
	return &Service{
		conn: db.SetupDB(),
	}
}

func (s *Service) LoginServiceHandler(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("No body has been posted.")
	}

	var records []db.UserRecord
	db.DecodeJSON(b, &records)

	u := records[0].Username
	p := records[0].Password

	err = db.QueryUserAndPassword(s.conn, u, p)

	res := db.LoginResponse{}

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		fmt.Println("error querying login user")
		res.Result = "fail"
		res.Message = "Incorrect username or password"
	} else {
		res.Result = "success"
		res.RedirectURL = "/home"
	}

	jsonData, err := json.Marshal(res)
	if err != nil {
		fmt.Println("failed marshalling data to json")
	}

	if err != nil {
		w.WriteHeader(401)
	} else {
		w.WriteHeader(303)
	}

	w.Write(jsonData)
}

func (s *Service) SignupServiceHandler(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("error reading request body")
	}

	res := db.UserSignupResponse{}

	var userRecord []db.UserRecord
	dec := json.NewDecoder(bytes.NewReader(b))
	dec.DisallowUnknownFields()
	err = dec.Decode(&userRecord)

	if err != nil {
		fmt.Printf("error decoding json body %v", err)
		res.GeneralMessage = "missing field or incorrect characters type in field"
		db.WriteJSONResponse(&res, w)
		return
	}

	e := userRecord[0].Email
	u := userRecord[0].Username
	p := userRecord[0].Password
	a := userRecord[0].Age

	if a < 1 {
		res.AgeError = "can not be smaller than 1"
		db.WriteJSONResponse(&res, w)
		return
	}

	if a < 18 {
		res.AgeError = "must be equal and greater than 18"
		db.WriteJSONResponse(&res, w)
		return
	}

	_, err = mail.ParseAddress(e)

	if err != nil {
		res.EmailError = "must be an email address"
		db.WriteJSONResponse(&res, w)
		return
	}

	var emailAndUserRecord db.UserRecord

	res = db.UserSignupResponse{}

	err = db.QueryEmailAndUser(s.conn, e, u, &emailAndUserRecord)

	if err != nil {
		res.EmailError = "this email or username has been registered"

		db.WriteJSONResponse(&res, w)
	} else {
		err = db.InsertUserForSignUp(s.conn, e, u, p, a)

		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			res.GeneralMessage = "signing up user failed for interal reason"
		} else {
			res.GeneralMessage = "signed up successfully"
			res.RedirectURL = "/home"
		}

		db.WriteJSONResponse(&res, w)
	}
}
