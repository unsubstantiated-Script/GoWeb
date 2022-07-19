package main

import (
	"WebGo/pkg/handlers"
	"fmt"
	"net/http"
)

const portNumber = ":8081"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("starting application on port: %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
