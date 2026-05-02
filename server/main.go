package main

import (
	"funcserver/server/db"
	"funcserver/server/page"
	"funcserver/server/service"
	"funcserver/server/session"
	"log"
	"net/http"
)

func main() {
	conn := db.SetupDB()

	mux := http.NewServeMux()

	// preflight

	// fetch session

	newSession := session.NewSessionManager()

	// router

	// serve content

	page.PagePipe(mux, newSession)

	// serve service
	service.ServicePipe(mux, conn, newSession)

	log.Fatal(http.ListenAndServe(":3000", mux))
}
