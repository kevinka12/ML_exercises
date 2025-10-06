package main

import (
	"fmt"
	"os"
	"strconv"
)

// Temperature struct
type Temperature struct {
	Value float64
	Unit  string
}

// Celsius → Fahrenheit
func celsiusToFahrenheit(t *Temperature) *Temperature {
	newValue := (t.Value * 9 / 5) + 32
	return &Temperature{Value: newValue, Unit: "F"}
}

// Fahrenheit → Celsius
func fahrenheitToCelsius(t *Temperature) *Temperature {
	newValue := (t.Value - 32) * 5 / 9
	return &Temperature{Value: newValue, Unit: "C"}
}

func main() {
	// Tjek at vi har nok arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage: tempconverter <value> <unit>")
		os.Exit(1)
	}

	// Læs argumenterne
	valueStr := os.Args[1] // fx "100"
	unit := os.Args[2]     // fx "C"

	// Konverter value fra string → float64
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		fmt.Println("Invalid temperature value. Please enter a numeric value.")
		os.Exit(1)
	}

	// Lav en Temperature struct
	temp := &Temperature{Value: value, Unit: unit}

	// Bestem hvilken konvertering vi skal lave
	var result *Temperature
	switch temp.Unit {
	case "C":
		result = celsiusToFahrenheit(temp)
	case "F":
		result = fahrenheitToCelsius(temp)
	default:
		fmt.Println("Invalid unit. Please provide 'C' for Celsius or 'F' for Fahrenheit.")
		os.Exit(1)
	}

	// Print resultatet (ingen decimaler)
	fmt.Printf("%.0f %s\n", result.Value, result.Unit)

}
