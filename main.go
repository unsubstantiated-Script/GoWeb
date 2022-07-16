package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8081"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("starting application on port: %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
