package main

import (
	server "funcserver/server/http"
	"log"
)

func main() {
	s := server.NewHTTPServer(":3000")

	log.Fatal(s.ListenAndServe())
}
