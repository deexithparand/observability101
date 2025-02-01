package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"observe/internal/server/response"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	var rf *response.ResponseFormat = &response.ResponseFormat{}

	var order_map map[string]string = map[string]string{
		"order 1": "greens",
		"order 2": "greens",
		"order 3": "greens",
		"order 4": "greens",
		"order 5": "greens",
	}

	orders_data, err := json.Marshal(order_map)
	if err != nil {
		panic(err)
	}

	err = rf.ResponseFormatSetter(w, r, string(orders_data))
	if err != nil {
		panic(err)
	}

	fmt.Println("Got the orders data")
}
