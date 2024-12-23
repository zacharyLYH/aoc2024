package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
	"zacharylyh/aoc2024"
)

func main() {
	file := aoc2024.ExtractTxtFile()
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sol := 0
	sol2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		rows := strings.Fields(line)

		sol += partOne(rows)
		sol2 += partTwo(rows)
	}

	aoc2024.HandleGenericErr(scanner.Err())

	fmt.Println("Day 2 part 1: ", sol)
	fmt.Println("Day 2 part 2: ", sol2)
}

func checker(prevVal, currVal, dir int) bool {
	if dir == 0 {
		return false
	}
	if !(math.Abs(float64(prevVal)-float64(currVal)) >= 1 && math.Abs(float64(prevVal)-float64(currVal)) <= 3) {
		return false
	}
	if dir == 1 && prevVal > currVal {
		return false
	}
	if dir == -1 && prevVal < currVal {
		return false
	}
	return true
}

func partTwo(rows []string) int {
	array := []int{}

	for _, row := range rows {
		num, err := strconv.Atoi(row)
		aoc2024.HandleGenericErr(err)
		array = append(array, num)
	}

	dir := 0
	if len(rows) >= 2 {
		if array[0] < array[1] {
			dir = 1
		} else if array[1] < array[0] {
			dir = -1
		}
	}
	for i:=1; i<len(array); i++ {
		complies := checker(array[i-1], array[i], dir)
		if !complies {
			removePrev := 0
			removeCurr := 0
			if i-1 >= 0 {
				prev := append(append([]string{}, rows[:i-1]...), rows[i:]...)
				removePrev = partOne(prev)
			}
			if i<len(rows) {
				curr := append(append([]string{}, rows[:i]...), rows[i+1:]...)
				removeCurr = partOne(curr)
			}
			if removeCurr == 1 || removePrev == 1 {
				return 1
			} else {
				return 0
			}
		}
	}
	return 1
}

func partOne(rows []string) int {
	prevVal := -1
	dir := 0
	isSafe := 1
	for _, row := range rows {
		val, err := strconv.Atoi(row)
		aoc2024.HandleGenericErr(err)
		if prevVal == -1 { //first value of the column
			prevVal = val
			continue
		}
		if dir == 0 { //set the direction of travel
			if prevVal < val {
				dir = 1
			} else if prevVal > val {
				dir = -1
			} else {
				isSafe = 0
				break
			}
		}
		if !(math.Abs(float64(prevVal)-float64(val)) >= 1 && math.Abs(float64(prevVal)-float64(val)) <= 3) {
			isSafe = 0
			break
		}
		if dir == 1 && prevVal > val {
			isSafe = 0
			break
		}
		if dir == -1 && prevVal < val {
			isSafe = 0
			break
		}
		prevVal = val
	}
	return isSafe
}