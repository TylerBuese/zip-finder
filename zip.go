package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func (m *Respository) zipfinder(w http.ResponseWriter, r *http.Request) {
	zip := r.URL.Query().Get("zip")
	rad := r.URL.Query().Get("radius")
	calcType := r.URL.Query().Get("type")
	w.Header().Set("Content-Type", "application/json")

	if zip == "" || rad == "" || calcType == "" {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Invalid query parameter. Query parameters: zip, radius, type must be provided.")
		return
	}

	radius, floatError := strconv.ParseFloat(rad, 64)
	if floatError != nil {
		fmt.Fprintf(w, "Unable to parse float ", rad, "Error:", floatError)
	}

	result, resultError := getAllZipsWithinRadius(zip, radius, calcType)
	if resultError != nil {
		fmt.Fprintf(w, resultError.Error())
	}

	jsonResult, jsonResultErr := json.Marshal(result)

	if jsonResultErr != nil {
		fmt.Fprintf(w, jsonResultErr.Error())
	}

	fmt.Fprintf(w, string(jsonResult))

}

func getAllZipsWithinRadius(zip string, rad float64, calcType string) ([]ZipCodes, error) {
	//Check for zip codes
	var selectedZip ZipCodes
	var foundZipCodes []ZipCodes
	for _, z := range zips {
		if z.Zip == zip {
			selectedZip = z
			//Zip code exists in database.
		}
	}

	if selectedZip.Zip == "" {
		return []ZipCodes{}, errors.New("The zip code " + zip + " does not exist in the database.")
	}

	for _, z := range zips {
		if CalculateDistance(selectedZip.Long, selectedZip.Lat, z.Long, z.Lat, calcType) < rad {
			foundZipCodes = append(foundZipCodes, z)
		}
	}

	return foundZipCodes, nil

}
