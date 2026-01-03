package main

import (
	"fmt"
	"errors"
	"math"
)

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

func main() {
	initial_prompts()
	val, base, dest, err := get_input()

	if err != nil {
		fmt.Println(err)
	} else {
		convert := float64(base) - float64(dest)
		result := val * math.Pow(10, convert)

		fmt.Println(result)
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
