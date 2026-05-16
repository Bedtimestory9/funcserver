package main

import (
	server "funcserver/server/http"
	"log"
)

func main() {
	// conn := db.SetupDB()

	s := server.NewHTTPServer(":3000")

	log.Fatal(s.ListenAndServe())

	// serve service
	// service.ServicePipe(mux, conn)
}
