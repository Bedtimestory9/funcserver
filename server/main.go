package main

import (
	"funcserver/server/db"
	"funcserver/server/page"
	"funcserver/server/service"
	"log"
	"net/http"
)

func main() {
	conn := db.SetupDB()

	mux := http.NewServeMux()

	mux.Handle("/styles.css", http.FileServer(http.Dir("public")))
	mux.Handle("/scripts/", http.FileServer(http.Dir("public")))

	// serve page content
	page.RouterPipe(mux)

	// serve service
	service.ServicePipe(mux, conn)

	log.Fatal(http.ListenAndServe(":3000", mux))
}
