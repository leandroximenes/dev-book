package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Listening port 3000")
	r := router.RunServer()

	log.Fatal(http.ListenAndServe(":3000", r))
}
