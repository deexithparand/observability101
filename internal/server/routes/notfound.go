package routes

import (
	"net/http"
	"observe/internal/server/response"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	var rf *response.ResponseFormat = &response.ResponseFormat{}

	err := rf.ResponseFormatSetter(w, r, "Page not found")
	if err != nil {
		panic(err)
	}
}
