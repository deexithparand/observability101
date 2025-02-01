package response

import (
	"fmt"
	"net/http"
)

type ResponseFormat struct {
	Method  string
	Route   string
	Message string
}

func (rf *ResponseFormat) ResponseFormatSetter(w http.ResponseWriter, r *http.Request, message string) error {
	// formatting the response
	response_setter := ResponseFormat{
		Method:  r.Method,
		Route:   r.RequestURI,
		Message: message,
	}

	response := fmt.Sprintf("Method : %s\nRoute : %s\nMessage %s\n", response_setter.Method, response_setter.Route, response_setter.Message)

	_, err := w.Write([]byte(response))
	return err
}
