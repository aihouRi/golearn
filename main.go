package main

import (
	"errors"
	"fmt"
	"net/http"
)

var portNumber = ":8080"

// Home is a home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About is a about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addVaules(2, 2)
	fmt.Fprintf(w, "This is the about page and 2 + 2 is %d", sum)
}

func addVaules(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideVaules(100.0, 10.0)
	if err != nil {
		fmt.Fprintf(w, "can not divide by 0")
		return
	}
	fmt.Fprintf(w, "%f divide by %f is %f", 100.0, 10.0, f)
}

func divideVaules(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("can not divide by zeor")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)


	fmt.Printf("Starting app on port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
