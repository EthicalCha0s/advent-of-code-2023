package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ExtractColours(text string) *ColourInput {
	colourInput := new(ColourInput)

	return colourInput
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

	for scanner.Scan() {
		line := scanner.Text()
		coloursText := strings.Split(line, ":")
		coloursText = strings.Split(strings.TrimSpace(coloursText[1]), ",")

		g := new(Game)
		for _, element := range coloursText {
			g.PassResults(ExtractColours(element))
		}
	}
}
