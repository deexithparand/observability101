package response

import (
	"fmt"
	"net/http"
)

type ResponseFormat struct {
	Method     string
	Route      string
	Message    string
	StatusCode int
}

func (rf *ResponseFormat) ResponseFormatSetter(w http.ResponseWriter, r *http.Request, message string, status_code int) error {

	// set status code in the Header
	w.WriteHeader(status_code)

	// formatting the response
	response_setter := ResponseFormat{
		Method:     r.Method,
		StatusCode: status_code,
		Route:      r.URL.Path,
		Message:    message,
	}

	response := fmt.Sprintf("StatusCode: %d\nMethod : %s\nRoute : %s\nMessage : %s\n", status_code, response_setter.Method, response_setter.Route, response_setter.Message)

	_, err := w.Write([]byte(response))
	return err
}
