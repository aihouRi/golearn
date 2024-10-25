package main

import (
	"fmt"
	"github.com/aihouRi/golearn/pkg/config"
	"github.com/aihouRi/golearn/pkg/handlers"
	"github.com/aihouRi/golearn/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.Appconfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("can not create template cache")
	}
	app.TemplateCache = tc

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting app on port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
