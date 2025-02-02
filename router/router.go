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

	// root route
	router.HandleFunc("/", handler.Root).Methods("GET")

	// user routes
	router.HandleFunc("/users", handler.GetUsers).Methods("GET")

	// orders routes
	router.HandleFunc("/orders", handler.GetOrders).Methods("GET")

	// Global 404 Handler
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		var path_not_found_rf *response.ResponseFormat
		err := path_not_found_rf.ResponseFormatSetter(w, r, "Path Not Found", 404)
		if err != nil {
			panic(err)
		}
	})

	return router
}
