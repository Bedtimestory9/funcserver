package service

import (
	"fmt"
	"funcserver/server/db"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func getUserMood(mux *http.ServeMux, conn *pgx.Conn) {
	mux.HandleFunc("GET /service/interaction/get-user-mood", func(w http.ResponseWriter, r *http.Request) {

		out, err := db.QueryUserMood(conn)
		if err != nil {
			fmt.Println("querying DB error")
		}

		w.Write([]byte(out))
	})
}
