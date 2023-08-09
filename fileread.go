package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadFile(path string) []ZipCodes {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	sc := bufio.NewScanner(file)
	//lines := make([]string, 0)

	var zips []ZipCodes

	i := 0
	for sc.Scan() {
		i++
		var entries = make([]byte, 0)
		inQuotes := false
		for _, text := range sc.Text() {
			if strings.Split(sc.Text(), ",")[0] == "65101" {

				if text == '"' && inQuotes {
					inQuotes = false
				}

				if text == '"' && !inQuotes {
					entries = append(entries, byte(text))
					inQuotes = true
				}
			}
		}
		zip := ZipCodes{
			Zip:    strings.Split(sc.Text(), ",")[0],
			State:  strings.Split(sc.Text(), ",")[8],
			City:   strings.Split(sc.Text(), ",")[3],
			County: strings.Split(sc.Text(), ",")[7],
		}
		if zip.Zip == "65101" {
			//log.Println(zip.County)
		}
		//lines = append(lines, sc.Text())
		zips = append(zips, zip)

	}

	return zips

}
