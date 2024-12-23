package aoc2024

import "fmt"

func HandleGenericErr(err error) {
	if err != nil {
		fmt.Println("Error message ", err)
	}
}