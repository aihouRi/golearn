package main

import (
	"fmt"
	"github.com/aihouRi/golearn/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"



func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting app on port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
