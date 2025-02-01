package routes

import "net/http"

var RoutePathMap map[string]string = map[string]string{
	// home page
	"Home": "/",

	// users
	"GetUsers": "/users",

	// orders
	"GetOrders": "/orders",
}

var RouteHandlerMap map[string]func(http.ResponseWriter, *http.Request) = map[string]func(http.ResponseWriter, *http.Request){
	// home page
	"/": GetHomePage,

	// users
	"/users": GetUsers,

	// orders
	"/orders": GetOrders,
}

func InitRoutes(mux *http.ServeMux) {
	for path, handler := range RouteHandlerMap {
		mux.HandleFunc(path, handler)
	}
}
