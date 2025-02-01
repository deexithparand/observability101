package server

import (
	"fmt"
	"net/http"
)

func GetHomePage(w http.ResponseWriter, r *http.Request) {

	response_setter := ResponseFormat{
		Method:  r.Method,
		Route:   r.RequestURI,
		Message: "This is the Home Page, can be used for Health check",
	}

	response := fmt.Sprintf("Method : %s\nRoute : %s\nMessage %s\n", response_setter.Method, response_setter.Route, response_setter.Message)

	w.Write([]byte(response))
}
