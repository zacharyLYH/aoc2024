package aoc2024

import (
	"fmt"
	"os"
)

func ExtractTxtFile() *os.File {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	return file
}