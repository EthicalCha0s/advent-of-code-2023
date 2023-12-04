package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "Hello32World"

	// Get the first digit
	firstDigit := str[0]
	if firstDigit >= '0' && firstDigit <= '9' {
		firstDigitString := string(firstDigit)
		firstDigitInt, err := strconv.Atoi(firstDigitString)
		if err == nil {
			fmt.Println("First digit:", firstDigitInt)
		}
	}

	// Get the last digit
	lastDigit := str[len(str)-1]
	if lastDigit >= '0' && lastDigit <= '9' {
		lastDigitString := string(lastDigit)
		lastDigitInt, err := strconv.Atoi(lastDigitString)
		if err == nil {
			fmt.Println("Last digit:", lastDigitInt)
		}
	}

	print(firstDigit)
	print(lastDigit)
}
