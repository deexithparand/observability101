package handler

import (
	"encoding/json"
	"net/http"
	"observe/response"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// orders data
	users := []User{
		{
			ID:   "1",
			Name: "Deexith",
		},
		{
			ID:   "2",
			Name: "Deexith",
		},
		{
			ID:   "3",
			Name: "Deexith",
		},
		{
			ID:   "4",
			Name: "Deexith",
		},
		{
			ID:   "5",
			Name: "Deexith",
		},
	}

	users_data_json, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}

	var get_users_rf *response.ResponseFormat
	err = get_users_rf.ResponseFormatSetter(w, r, string(users_data_json), 200)
	if err != nil {
		panic(err)
	}
}
