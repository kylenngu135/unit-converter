package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
)

type ConversionRequest struct {
    Value       float64 `json:"value"`
    CurrentUnit string  `json:"current"`
    ConvertUnit string  `json:"convert"`
}

type ConversionResponse struct {
    Result float64 `json:"result"`
    Error  string  `json:"error,omitempty"`
}

var conversionFactors = map[string]float64{
    "millimeter": 0.001,
    "centimeter": 0.01,
    "decimeter":  0.1,
    "meter":      1.0,
    "dekameter":  10.0,
    "hectometer": 100.0,
    "kilometer":  1000.0,
}

func convert(value float64, from string, to string) (float64, error) {
	fromFactor, ok1 := conversionFactors[from]
	toFactor, ok2 := conversionFactors[to]


	if !ok1 || !ok2 {
		return 0, fmt.Errorf("invalid unit")
	}

	meters := value * fromFactor
	result := meters / toFactor

	fmt.Println("value:", value)
	fmt.Println("fromFactor: ", fromFactor)
	fmt.Println("toFactor: ", toFactor)
	fmt.Println("Result: ", result)

	return result, nil
}

func handleConvert(w http.ResponseWriter, r *http.Request) {
    // Set CORS headers to allow requests from different origins (like local HTML files)
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    w.Header().Set("Content-Type", "application/json") // Tell the client we're sending JSON

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ConversionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		json.NewEncoder(w).Encode(ConversionResponse{Error: "Invalid request"})
		return
	}

	result, err := convert(req.Value, req.CurrentUnit, req.ConvertUnit)

	if err != nil {
		json.NewEncoder(w).Encode(ConversionResponse{Error: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(ConversionResponse{Result: result})
}

func main() {
	http.HandleFunc("/convert", handleConvert)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
