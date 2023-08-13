package main

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

}
