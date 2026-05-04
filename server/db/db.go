// Package db deals with db
package db

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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

func WriteJSONResponse[R Response](res *R, w http.ResponseWriter) {
	w.WriteHeader(401)
	jsonData, err := json.Marshal(res)
	if err != nil {
		fmt.Println("failed marshalling data to json")
	}
	w.Write(jsonData)
}

func QueryUserMood(conn *pgx.Conn, record *MoodRecord) error {
	err := conn.QueryRow(context.Background(),
		"SELECT name, mood FROM golang_table WHERE name='lawrence'").
		Scan(&record.Name, &record.Mood)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}
	return err
}

func QueryLoginUser(conn *pgx.Conn, u string, p string) error {
	var record UserRecord
	err := conn.QueryRow(context.Background(),
		"SELECT username, password FROM users WHERE username=$1 AND password=$2", u, p).
		Scan(&record.Username, &record.Password)
	return err
}

func QuerySignupUser(conn *pgx.Conn, e string, u string, record *UserRecord) error {
	emailErr := conn.QueryRow(context.Background(),
		"SELECT email FROM users WHERE email=$1", e).
		Scan(&record.Email)

	usernameErr := conn.QueryRow(context.Background(),
		"SELECT username FROM users WHERE username=$1", u).
		Scan(&record.Username)

	// != nil meaning can't find it
	if emailErr != nil && usernameErr != nil {
		fmt.Println("user hasn't been registered")
		return nil
	} else {
		err := errors.New("user has been registered")
		return err
	}
}

func InsertSignupUser(conn *pgx.Conn, e string, u string, p string, a int) error {
	_, err := conn.Exec(context.Background(),
		"INSERT INTO users (email, username, password, age) VALUES ($1, $2, $3, $4)",
		e, u, p, a,
	)

	return err
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
