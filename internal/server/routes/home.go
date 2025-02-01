package routes

import (
	"fmt"
	"net/http"
	"observe/internal/server/response"
)

func GetHomePage(w http.ResponseWriter, r *http.Request) {

	var rf *response.ResponseFormat = &response.ResponseFormat{}

	err := rf.ResponseFormatSetter(w, r, "This is the Home Page, can be used for Health check")
	if err != nil {
		panic(err)
	}

	fmt.Println("Got the home page data")
}
