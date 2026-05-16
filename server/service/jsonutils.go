package service

import (
	"encoding/json"
	"fmt"
	"funcserver/server/db"
	"net/http"
)

func WriteJSONResponse[R db.Response](res *R, w http.ResponseWriter) {
	w.WriteHeader(401)
	jsonData, err := json.Marshal(res)
	if err != nil {
		fmt.Println("failed marshalling data to json")
	}
	w.Write(jsonData)
}
