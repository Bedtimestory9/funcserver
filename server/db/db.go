// Package db deals with db
package db

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"os"

	"github.com/jackc/pgx/v5"
)

func EncodeJSON(record any) string {
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	err := enc.Encode(record)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Encoding json error: %v\n", err)
	}
	out := b.String()
	return out
}

func DecodeJSON[R Record](jsonData []byte, records *[]R) *[]R {
	err := json.Unmarshal(jsonData, records)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unmarshalling json error: %v\n", err)
	}
	return records
}

func QueryUserMood(conn *pgx.Conn) (string, error) {
	var record MoodRecord
	err := conn.QueryRow(context.Background(),
		"SELECT name, mood FROM golang_table WHERE name='lawrence'").
		Scan(&record.Name, &record.Mood)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	return EncodeJSON(record), err
}

func QueryLoginUser(conn *pgx.Conn, u string, p string) (LoginResponse, error) {
	var record UserRecord
	err := conn.QueryRow(context.Background(),
		"SELECT username, password FROM users WHERE username=$1 AND password=$2", u, p).
		Scan(&record.Username, &record.Password)
	res := LoginResponse{}

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		res.Result = "fail"
		res.Message = "Incorrect username or password"
		return res, err
	}

	res.Result = "success"
	res.RedirectURL = "/home/" + u

	return res, nil
}

func QuerySignupUser(conn *pgx.Conn, e string, u string) (UserSignupResponse, bool) {
	var record UserSignupRecord
	res := UserSignupResponse{}

	emailErr := conn.QueryRow(context.Background(),
		"SELECT email FROM users WHERE email=$1", e).
		Scan(&record.Email)

	usernameErr := conn.QueryRow(context.Background(),
		"SELECT username FROM users WHERE username=$1", u).
		Scan(&record.Username)

	// != nil meaning can't find it
	if emailErr != nil && usernameErr != nil {
		fmt.Fprintf(os.Stderr, "Email QueryRow failed: %v, username QueryRow failed: %v but it's intended\n", emailErr, usernameErr)
		return res, true
	} else {
		res.EmailError = "This emails or username has been registered"
		return res, false
	}
}

func InsertSignupUser(conn *pgx.Conn, e string, u string, p string, a int, res *UserSignupResponse) UserSignupResponse {
	_, err := conn.Exec(context.Background(),
		"INSERT INTO users (email, username, password, age) VALUES ($1, $2, $3, $4)",
		e, u, p, a,
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		res.GeneralMessage = "Signing up user failed for interal reason"
	} else {
		res.GeneralMessage = "Signed up successfully"
		res.RedirectURL = "/home/" + u
	}

	return *res
}

func SetupDB() *pgx.Conn {
	readEnv, err := os.ReadFile(".env")
	connString := strings.TrimSpace(string(readEnv))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read file")
	}
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}
