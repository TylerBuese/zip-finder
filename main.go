package main

import (
	"log"
	"net/http"
)

type ZipCodes struct {
	Zip    string
	City   string
	State  string
	Lat    float64
	Long   float64
	County string
}

var path = "./zip_code_database.csv"
var zips = ReadFile(path)

const port = ":8080"

func main() {
	app := AppConfig{
		TemplateCache: nil,
		UseCache:      false,
	}

	NewTemplates(&app)

	repo := NewRepo(&app)
	NewHandlers(repo)

	serv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err := serv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
