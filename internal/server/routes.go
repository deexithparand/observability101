package server

import "net/http"

var RoutePathMap map[string]string = map[string]string{
	// home page
	"home": "/",

	// users
	// "users": "/users",

	// orders
	// "orders": "/orders",
}

var RouteHandlerMap map[string]func(http.ResponseWriter, *http.Request) = map[string]func(http.ResponseWriter, *http.Request){
	// home page
	"/": GetHomePage,

	// users

	// orders
}
