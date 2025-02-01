package main

import (
	"fmt"
	"net/http"
	"observe/internal/server/routes"
	"strconv"
)

func server_closed() {
	fmt.Println("End of Server")
}

func main() {

	// server config
	host := "localhost:"
	port := 8080

	// defining mux
	mux := http.NewServeMux()

	// defining routes
	routes.InitRoutes(mux)

	// start server
	httpserver := &http.Server{
		Addr: host + strconv.Itoa(port),
	}

	fmt.Println("Server's heart beats at : ", httpserver.Addr)
	err := http.ListenAndServe(httpserver.Addr, mux)
	if err != nil {
		httpserver.ErrorLog.Panicf("server error : %v", err)
	}

	defer server_closed()

}
