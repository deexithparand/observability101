package router

import (
	"net/http"
	handler "observe/handlers"
	"observe/response"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	// router & defined routes
	router.HandleFunc("/", handler.Root).Methods("GET")

	// router.HandleFunc("/users", handl)

	// Global 404 Handler
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		var path_not_found_rf *response.ResponseFormat
		err := path_not_found_rf.ResponseFormatSetter(w, r, "Path Not Found")
		if err != nil {
			panic(err)
		}
	})

	return router
}
