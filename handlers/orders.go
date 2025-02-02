package handler

import (
	"encoding/json"
	"net/http"
	"observe/response"
)

type Order struct {
	ID     string `json:"id"`
	Amount int    `json:"amount"`
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	// orders data
	orders := []Order{
		{
			ID:     "1",
			Amount: 100,
		},
		{
			ID:     "2",
			Amount: 200,
		},
		{
			ID:     "3",
			Amount: 300,
		},
		{
			ID:     "4",
			Amount: 400,
		},
		{
			ID:     "5",
			Amount: 500,
		},
	}

	orders_data_json, err := json.Marshal(orders)
	if err != nil {
		panic(err)
	}

	var get_orders_rf *response.ResponseFormat
	err = get_orders_rf.ResponseFormatSetter(w, r, string(orders_data_json), 200)
	if err != nil {
		panic(err)
	}
}
