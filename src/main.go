package main

import (
	"errors"
	"fmt"
	"math"
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

	return result, nil
}

func handleConvert(w http.ResponseWriter, r *http.Request) {

}

func main() {

}

/*
type Factor int

const (
	Mm Factor = iota
	Cm
	Dm
	M
	Dam
	Hm
	Km
)

func (f Factor) String() string {
	switch f {
	case Mm:
		return "Mm"
	case Cm:
		return "Cm"
	case Dm:
		return "Dm"
	case M:
		return "M"
	case Dam:
		return "Dam"
	case Hm:
		return "Hm"
	case Km:
		return "Km"
	default:
		return "unknown"
	}
}

func str_to_factor(f string) (factor Factor, err error) {
	switch f {
	case "Mm":
		return Mm, nil
	case "Cm":
		return Cm, nil
	case "Dm":
		return Dm, nil
	case "M":
		return M, nil
	case "Dam":
		return Dam, nil
	case "Hm":
		return Hm, nil
	case "Km":
		return Km, nil
	default:
		return M, errors.New("invalid input")
	}
}

func initial_prompts() {
	fmt.Println("This is the CLI version of the unit-converter program.")
}

func get_input() (val float64, base Factor, dest Factor, err error) {
	var base_str string
	var dest_str string
	var base_err error
	var dest_err error

	fmt.Print("Enter the base value: ")
	fmt.Scanf("%f", &val)
	fmt.Print("Enter the base unit: ")
	fmt.Scanf("%s", &base_str)
	fmt.Print("Enter the new unit: ")
	fmt.Scanf("%s", &dest_str)

	base, base_err = str_to_factor(base_str)
	dest, dest_err = str_to_factor(dest_str)

	if base_err != nil {
		err = fmt.Errorf("base factor: %s", base_err)
	}

	if dest_err != nil {
		if base_err != nil {
			err = fmt.Errorf("%w, dest factor: %s", err, dest_err)
		} else {
			err = fmt.Errorf("dest factor: %s", dest_err)
		}
	}

	return
}

func main() {
	//initial_prompts()
	// val, base, dest, err := get_input()

	if err != nil {
		fmt.Println(err)
	} else {
		convert := float64(base) - float64(dest)
		result := val * math.Pow(10, convert)

		fmt.Println(result)
	}
}
*/
