package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Define a dictionary for text-to-digit mappings
var textToDigitMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getKeys(dict map[string]int) []string {
	keys := []string{}
	for key := range dict {
		keys = append(keys, key)
	}
	return (keys)
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func stringToInt(text string) int {
	textDigits := getKeys(textToDigitMap)
	if stringInSlice(text, textDigits) {
		return textToDigitMap[text]
	}

	i, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return (i)
}

func extractDigits(text string) int {
	textDigits := getKeys(textToDigitMap)
	regexPattern := `((\d)|` + strings.Join(textDigits, "|") + ")"
	digitsRegex := regexp.MustCompile(regexPattern)
	digits := digitsRegex.FindAllString(text, -1)

	return (stringToInt(digits[0]) * 10) + stringToInt(digits[len(digits)-1])
}

func main() {
	// Open connection to text-file
	filename := "input2.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("\n" + line)
		fmt.Println(extractDigits(line))

		sum += extractDigits((line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println(sum)
}
