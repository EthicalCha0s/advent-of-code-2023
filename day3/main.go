package main

import (
	"bufio"
	"fmt"
	"os"
)

func GetFileLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func main() {
	// Open connection to text-file
	lines := GetFileLines("input.txt")
	fmt.Println(lines)

}
