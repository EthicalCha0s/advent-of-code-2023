package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// // Define a dictionary for text-to-digit mappings
// var textToDigitMap = map[string]int{
// 	"one":   1,
// 	"two":   2,
// 	"three": 3,
// 	"four":  4,
// 	"five":  5,
// 	"six":   6,
// 	"seven": 7,
// 	"eight": 8,
// 	"nine":  9,
// }

func stringToInt(text string) int {
	i, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return (i)
}

func extractDigits(text string) int {
	digitsRegex := regexp.MustCompile(`\d`)
	digits := digitsRegex.FindAllString(text, -1)

	if len(digits) == 1 {
		return stringToInt(digits[0] + digits[0])
	} else {
		return stringToInt(digits[0] + digits[len(digits)-1])
	}
}

func main() {
	// Open connection to text-file
	filename := "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		sum += extractDigits((line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println(sum)
}
