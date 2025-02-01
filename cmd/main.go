package main

import (
	"fmt"
	"net/http"
	"observe/internal/server"
	"strconv"
)

func server_closed() {
	fmt.Println("End of Server")
}

func main() {

	// server config
	host := "localhost:"
	port := 8080

	// defining routes
	for route_name, route_path := range server.RoutePathMap {
		fmt.Println("Defined route for ", route_name)
		http.HandleFunc(route_path, server.RouteHandlerMap[route_path])
	}

	// start server
	server := &http.Server{
		Addr: host + strconv.Itoa(port),
	}

	fmt.Println("Server's heart beats at : ", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		server.ErrorLog.Panicf("server error : %v", err)
	}

	defer server_closed()

}
