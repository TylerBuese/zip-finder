package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	path := "./zip_code_database.csv"
	zips := ReadFile(path)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		text, marshalErr := json.Marshal(zips)
		if marshalErr != nil {
			log.Println(marshalErr)
		}

		_, err := fmt.Fprintf(w, string(text))
		log.Println(r.Method)

		if err != nil {

		}
	})

	_ = http.ListenAndServe(":8080", nil)

}
