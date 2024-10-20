package main

import (
	"net/http"
)

// Home is a home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

// About is a about page handler
func About(w http.ResponseWriter, r *http.Request) {

}
