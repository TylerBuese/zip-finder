package main

import (
	"log"
	"math"
	"strings"
)

func ToRadians(angle float64) float64 {
	return (float64(angle) * math.Pi) / 180
}

func Distance(lon1, lon2, lat1, lat2 float64) float64 {
	lon1 = ToRadians(lon1)
	lon2 = ToRadians(lon2)
	lat1 = ToRadians(lat1)
	lat2 = ToRadians(lat2)

	dlon := lon2 - lon1
	dlat := lat2 - lat1

	a := math.Pow(math.Sin(dlat/2), 2) +
		math.Cos(lat1)*
			math.Cos(lat2)*math.Pow(math.Sin(dlon/2), 2)
	c := 2 * math.Asin(math.Sqrt(a))
	r := float64(6371)

	return c * r
}

func CalculateDistance(lon1, lat1, lon2, lat2 float64, calcType string) float64 {
	if strings.ToLower(calcType) == "mi" {
		return Distance(lon1, lon2, lat1, lat2) / 1.609
	} else if strings.ToLower(calcType) == "km" {
		return Distance(lon1, lon2, lat1, lat2)
	} else {
		log.Fatal("Invalid calcualtion type provided. Provided type was: ", calcType)
	}

	return float64(0)

}
