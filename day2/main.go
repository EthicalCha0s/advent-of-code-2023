package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GetNumber(text string) int {
	re := regexp.MustCompile(`\d+`)
	match := re.FindString(text)

	if match != "" {
		number, _ := strconv.Atoi(match)
		return number
	} else {
		return 0
	}
}

func ExtractColourInput(text string) *ColourInput {
	coloursTextList := strings.Split(text, ",")

	coloursMap := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, element := range coloursTextList {
		element = strings.TrimSpace(element)
		splitElement := strings.Split(element, " ")

		amount, _ := strconv.Atoi(splitElement[0]) // GetNumber(element)
		colourString := splitElement[1]

		coloursMap[colourString] = amount
	}

	colourInput := new(ColourInput)
	colourInput.R = coloursMap["red"]
	colourInput.G = coloursMap["green"]
	colourInput.B = coloursMap["blue"]

	return colourInput
}

func SumGameIDs(games []Game) int {
	idSum := 0
	for _, game := range games {
		idSum += game.ID
	}

	return idSum
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

	possibleGames := []Game{}

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, ":")
		coloursText := strings.Split(strings.TrimSpace(splitLine[1]), ";")

		g := new(Game)
		g.ID = GetNumber(splitLine[0])
		for _, element := range coloursText {
			g.PassResults(ExtractColourInput(element))
		}

		testInput := &ColourInput{
			R: 12,
			G: 13,
			B: 14,
		}

		if g.IsPossibleWithCubes(testInput) {
			possibleGames = append(possibleGames, *g)
		}
	}

	fmt.Println("The answer is:", SumGameIDs(possibleGames))
}
