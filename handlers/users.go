package handler

import "net/http"

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

}
