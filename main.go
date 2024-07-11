package main

import (
	"fmt"
	"log"
	"myapi/router"
	"net/http"
)

func main() {
	r := router.NewRouter()
	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
