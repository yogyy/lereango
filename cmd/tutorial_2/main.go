package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = `Hello World`
	PrintMe(printValue)

	numerator := 10
	denominator := 0
	var res, remainder, err = intDivision(numerator, denominator)

	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v", res)
	default:
		fmt.Printf("The result of the integer division is %v with remainder %v", res, remainder)
	}

	switch remainder {
	case 0:
		fmt.Println("The division was exact")
	case 1, 2:
		fmt.Println("The division was close")
	default:
		fmt.Println("The divison was no close")
	}
}

func PrintMe(value string) {
	fmt.Println(value)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var res int = numerator / denominator
	var remainder int = numerator % denominator

	return res, remainder, err
}
