package main

import (
	"fmt"
	"log"
	"main/src/router"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	r := router.RunServer()

	fmt.Printf("Listening port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
