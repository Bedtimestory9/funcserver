package page

import (
	"encoding/json"
	"fmt"
	"funcserver/server/db"
	"io"
	"net/http"
	"os"
)

type Login struct {
}

func (l *Login) PostLoginHandler(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("No body has been posted.")
	}

	var records []db.UserRecord
	db.DecodeJSON(b, &records)

	u := records[0].Username
	p := records[0].Password

	conn := db.SetupDB()

	err = db.QueryUserAndPassword(conn, u, p)

	res := db.LoginResponse{}

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		fmt.Println("error querying login user")
		res.Result = "fail"
		res.Message = "Incorrect username or password"
	} else {
		res.Result = "success"
		res.RedirectURL = "/home"
	}

	jsonData, err := json.Marshal(res)
	if err != nil {
		fmt.Println("failed marshalling data to json")
	}

	if err != nil {
		w.WriteHeader(401)
	} else {
		w.WriteHeader(303)
	}

	w.Write(jsonData)
}
