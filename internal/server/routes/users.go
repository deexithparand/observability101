package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"observe/internal/server/response"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var rf *response.ResponseFormat = &response.ResponseFormat{}

	var user_map map[string]string = map[string]string{
		"user 1": "people",
		"user 2": "people",
		"user 3": "people",
		"user 4": "people",
		"user 5": "people",
	}

	users_data, err := json.Marshal(user_map)
	if err != nil {
		panic(err)
	}

	err = rf.ResponseFormatSetter(w, r, string(users_data))
	if err != nil {
		panic(err)
	}

	fmt.Println("Got the users data")
}
