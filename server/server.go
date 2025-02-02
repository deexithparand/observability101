package server

import (
	"fmt"
	"net/http"
	"observe/config"
	"observe/router"
)

func StartServer() {
	// server config
	sc := config.InitServerConfig()
	addr := sc.Host + ":" + sc.Port

	// router init
	router := router.InitRouter()

	// server start
	fmt.Println("Server Running at ", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
