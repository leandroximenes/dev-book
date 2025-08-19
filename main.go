package main

import (
	"fmt"
	"log"
	"main/src/config"
	"main/src/router"
	"net/http"
)

func main() {
	config.Load()
	fmt.Printf("Listening Port %d\n", config.Port)

	addr := fmt.Sprintf(":%d", config.Port)
	r := router.RunServer()
	log.Fatal(http.ListenAndServe(addr, r))
}
