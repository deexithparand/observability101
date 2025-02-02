package handler

import (
	"log"
	"net/http"
	"observe/response"
)

func Root(w http.ResponseWriter, r *http.Request) {
	var root_rf *response.ResponseFormat
	err := root_rf.ResponseFormatSetter(w, r, "Welcome to the Observe API", 200)
	if err != nil {
		panic(err)
	}
	log.Println("Response Sent : Root URL")
}
