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

type Record interface {
	MoodRecord | UserRecord
}

type Response interface {
	ValidationResponse
}

type MoodRecord struct {
	Name string `json:"name"`
	Mood int    `json:"mood"`
}

type UserRecord struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ValidationResponse struct {
	Result  string `json:"result"`
	Message string `json:"message"`
}

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

func QueryUser(conn *pgx.Conn, u string, p string) error {
	var record UserRecord
	err := conn.QueryRow(context.Background(),
		"SELECT username, password FROM users WHERE username='"+u+"' AND password='"+p+"'").
		Scan(&record.Username, &record.Password)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

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
