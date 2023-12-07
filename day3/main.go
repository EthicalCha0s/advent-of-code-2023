package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

const ROWS = 140
const COLUMNS = 140

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

	twodim_array := [ROWS][COLUMNS]rune{}

	for index, line := range lines {
		for count, runeval := range line {
			twodim_array[index][count] = runeval
		}
	}

	current_num := ""
	running_count := 0
	hasSurrounded := false

	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLUMNS; j++ {

			if unicode.IsDigit(twodim_array[i][j]) {
				current_num = current_num + string(twodim_array[i][j])
				// if surrounded by something, set the flag
				if isSurrounded(twodim_array, i, j) {
					hasSurrounded = true
				}
			}

			if !(unicode.IsDigit(twodim_array[i][j])) {
				// number finished - if we've processed the number, and it is surrounded by anything, add to total
				if hasSurrounded {
					intval, _ := strconv.Atoi(current_num)
					running_count = running_count + intval
					hasSurrounded = false
				}
				current_num = ""
			}
		}
	}
	fmt.Println(running_count)
}

func isSurrounded(arr [ROWS][COLUMNS]rune, row int, col int) bool {
	// if top left exists
	if (row-1 != -1) && (col-1 != -1) {
		if !unicode.IsDigit(arr[row-1][col-1]) && arr[row-1][col-1] != '.' {
			return true
		}
	}
	if row-1 != -1 {
		// follows that row-1, col will exist
		if !unicode.IsDigit(arr[row-1][col]) && arr[row-1][col] != '.' {
			return true
		}
	}
	// now check if row-1 exists
	if col-1 != -1 {
		if !unicode.IsDigit(arr[row][col-1]) && arr[row][col-1] != '.' {
			return true
		}
	}
	// now check if top right exists
	if (row-1 != -1) && (col+1 != COLUMNS) {
		if !unicode.IsDigit(arr[row-1][col+1]) && arr[row-1][col+1] != '.' {
			return true
		}
	}
	// now check if r+1, col-1 exists
	if (row+1 != ROWS) && (col-1 != -1) {
		if !unicode.IsDigit(arr[row+1][col-1]) && arr[row+1][col-1] != '.' {
			return true
		}
	}
	// now check if r+1, col exists
	if row+1 != ROWS {
		if !unicode.IsDigit(arr[row+1][col]) && arr[row+1][col] != '.' {
			return true
		}
	}

	// now check if r+1, c+1 exists
	if (row+1 != ROWS) && (col+1 != COLUMNS) {
		if !unicode.IsDigit(arr[row+1][col+1]) && arr[row+1][col+1] != '.' {
			return true
		}
	}

	if col+1 != COLUMNS {
		if !unicode.IsDigit(arr[row][col+1]) && arr[row][col+1] != '.' {
			return true
		}
	}
	return false
}
