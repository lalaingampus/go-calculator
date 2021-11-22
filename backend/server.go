package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Numbers represents numebers model
type numbers struct {
	Num1 float64 `json:"num1"`
	Num2 float64 `json:"num2"`
}

// Numbers represents numebers model
type numsResponseData struct {
	Add float64 `json:"add"`
	Mul float64 `json:"mul"`
	Sub float64 `json:"sub"`
	Div float64 `json:"div"`
}

// Process functions are used to calculations
func process(numsdata numbers) numsResponseData {
	var numsres numsResponseData
	numsres.Add = numsdata.Num1 + numsdata.Num2
	numsres.Mul = numsdata.Num1 * numsdata.Num2
	numsres.Sub = numsdata.Num1 - numsdata.Num2
	numsres.Div = numsdata.Num1 / numsdata.Num2

	return numsres
}

// For write function which will be called to process JSON comin data to backend
func calc(w http.ResponseWriter, request *http.Request) {
	// Decoder to decode the incoming JSON ata from request body
	decode := json.NewDecoder(request.Body)

	var numsData numbers
	var numsResData numsResponseData

	decode.Decode(&numsData)

	numsResData = process(numsData)

	fmt.Println(numsResData)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	// for check the error from responsewriter
	if err := json.NewEncoder(w).Encode(numsResData); err != nil {
		panic(err)
	}
}

// For definition HTTP routing.
func main() {
	http.HandleFunc("/", calc)
	http.ListenAndServe(":8090", nil)
}
