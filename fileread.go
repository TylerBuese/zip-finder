package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func ReadFile(path string) []ZipCodes {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var zips []ZipCodes

	r := csv.NewReader(file)

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		lat, _ := strconv.ParseFloat(record[12], 64)
		long, _ := strconv.ParseFloat(record[13], 64)

		zip := ZipCodes{
			Zip:    record[0],
			State:  record[6],
			Lat:    lat,
			Long:   long,
			City:   record[3],
			County: record[7],
		}

		zips = append(zips, zip)
	}

	return zips

}
