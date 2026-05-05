package service

import (
	"fmt"
	"funcserver/server/db"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func getUserMood(mux *http.ServeMux, conn *pgx.Conn) {
	mux.HandleFunc("GET /service/interaction/get-user-mood", func(w http.ResponseWriter, r *http.Request) {
		var record db.MoodRecord
		err := db.QueryUserAndMood(conn, &record)

		if err != nil {
			fmt.Println("querying DB error")
		}

		out := db.EncodeJSON(record)
		w.Write([]byte(out))
	})
}
